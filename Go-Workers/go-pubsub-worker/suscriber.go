package main

import (
	"context"
	encoder "encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"cloud.google.com/go/pubsub"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	credentials_path := os.Getenv("PUBSUB_KEY_PATH")
	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", credentials_path)
	projectID := os.Getenv("PUBSUB_PROJECT")
	subID := os.Getenv("GOLANG_SUB")

	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	defer client.Close()

	fmt.Println("Esperando mensajes...")

	sub := client.Subscription(subID)
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	cm := make(chan *pubsub.Message)
	defer close(cm)

	go func() {
		for msg := range cm {

			/*GUARDAR INFO EN REDIS*/
			cadena := msg.Data
			atributos := msg.Attributes
			m := make(map[string]string)
			m["mensajePubSub"] = string(cadena)
			for k, v := range atributos {
				m[k] = v
			}

			data, err := encoder.Marshal(m)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
			}

			redisHOST := os.Getenv("REDIS_HOST")
			redisPORT := os.Getenv("REDIS_PORT")
			redisURL := fmt.Sprintf("%v:%v", redisHOST, redisPORT)

			ctxRedis := context.Background()

			cliente := redis.NewClient(&redis.Options{
				Addr:     redisURL,
				Password: "", // no password set
				DB:       0,  // use default DB
			})

			randomKey := strings.Replace(uuid.New().String(), "-", "", -1)

			err = cliente.Set(ctxRedis, randomKey, data, 0).Err()
			if err != nil {
				panic(err)
			}

			fmt.Println("Mensaje recibido :D")
			fmt.Printf("Data: %q\n", string(cadena))
			fmt.Println("Attributes:")
			for key, value := range atributos {
				fmt.Printf("%s = %s\n", key, value)
			}
			msg.Ack()
		}
	}()

	err = sub.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
		cm <- msg
	})
	if err != nil {
		fmt.Printf("Receive: %v\n", err)
	}

}
