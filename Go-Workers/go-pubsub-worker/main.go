package main

import (
	"context"
	encoder "encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"cloud.google.com/go/pubsub"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
			m["service"] = "Google PubSub"
			for k, v := range atributos {
				m[k] = v
			}

			fmt.Println(">> PubSub: Mensaje recibido :D")
			fmt.Printf("Data: %q\n", string(cadena))
			fmt.Println("Attributes:")
			for key, value := range atributos {
				fmt.Printf("%s = %s\n", key, value)
			}

			data, err := encoder.Marshal(m)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
			}

			redisHOST := os.Getenv("REDIS_HOST")
			redisPORT := os.Getenv("REDIS_PORT")
			redisPASS := os.Getenv("REDIS_PASS")
			redisURL := fmt.Sprintf("%v:%v", redisHOST, redisPORT)

			ctxRedis := context.Background()

			cliente := redis.NewClient(&redis.Options{
				Addr:     redisURL,
				Password: redisPASS, // no password set
				DB:       0,         // use default DB
			})

			randomKey := strings.Replace(uuid.New().String(), "-", "", -1)

			err = cliente.Set(ctxRedis, randomKey, data, 0).Err()
			if err != nil {
				panic(err)
			} else {
				fmt.Println(">> PubSub: Mensaje guardado en Redis :)")
			}

			/*GUARDAR LOGS EN MONGO*/

			mongoHOST := os.Getenv("MONGO_HOST")
			mongoPORT := os.Getenv("MONGO_PORT")
			mongoDB := os.Getenv("MONGO_DB")
			mongoCOL := os.Getenv("MONGO_COL")
			mongoUSER := os.Getenv("MONGO_USER")
			mongoPASS := os.Getenv("MONGO_PASS")
			mongoURL := fmt.Sprintf(`mongodb://%v:%v/?authSource=admin&readPreference=primary&directConnection=true&ssl=false`, mongoHOST, mongoPORT)

			credential := options.Credential{
				Username: mongoUSER,
				Password: mongoPASS,
			}

			ctxMongo, cancel := context.WithTimeout(context.Background(), time.Second*10)
			clientOptions := options.Client().ApplyURI(mongoURL).SetAuth(credential)

			c, err := mongo.NewClient(clientOptions)
			if err != nil {
				log.Fatalf("Error al crear cliente %v", err)
			}
			err = c.Connect(ctxMongo)
			if err != nil {
				log.Fatalf("Error al realizar conexion %v", err)
			}

			err = c.Ping(ctxMongo, nil)
			if err != nil {
				log.Fatalf("Error al conectar %v", err)
			}

			ctxInsert := context.Background()

			m2 := make(map[string]string)
			m2["service"] = "Google PubSub"
			for k, v := range atributos {
				m2[k] = v
			}

			todoCollection := c.Database(mongoDB).Collection(mongoCOL)
			r, err := todoCollection.InsertOne(ctxInsert, m2)
			if err != nil {
				log.Fatalf("Error al guardar logs %v", err)
			} else {
				fmt.Println("PubSub >> Log guardado en Mongo con el id: ", r.InsertedID)
			}

			/*-----------------------------------------------------*/
			c.Disconnect(ctxInsert)
			cancel()
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
