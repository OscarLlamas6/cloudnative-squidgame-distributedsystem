package main

import (
	"context"
	encoder "encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/segmentio/kafka-go"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	ctxKafKa := context.Background()

	kafkaHOST := os.Getenv("KAFKA_HOST")
	kafkaPORT := os.Getenv("KAFKA_PORT")
	kafkaTOPIC := os.Getenv("KAFKA_TOPIC")
	brokerURL := fmt.Sprintf("%v:%v", kafkaHOST, kafkaPORT)

	//l := log.New(os.Stdout, ">> KAFKA SUSCRIBER: ", 0)

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{brokerURL},
		Topic:   kafkaTOPIC,
		GroupID: "squidgames",
		//Logger:      l,
		StartOffset: kafka.LastOffset,
		MinBytes:    0,
		MaxBytes:    10e6, //10MB
	})

	fmt.Println("Esperando mensajes...")

	for {
		//Consumiento mensajes del Topic
		msg, err := r.ReadMessage(ctxKafKa)
		if err != nil {
			panic("could not read message " + err.Error())
		} else {
			// after receiving the message, log its value
			fmt.Println("Kafka >> Mensaje recibido :D")
			//fmt.Println(string(msg.Value))

			/*GUARDAR EN REDIS*/

			//Creo un map para mappea el arreglo de bytes del cuerpo del mensaje
			m := make(map[string]string)
			err2 := encoder.Unmarshal(msg.Value, &m)
			if err2 != nil {
				fmt.Printf("Error: %v\n", err)
			}
			//Itero en cada llave/valor del map
			for key, value := range m {
				fmt.Printf("%s = %s\n", key, value)
			}

			/*GUARDAR EN REDIS*/
			redisHOST := os.Getenv("REDIS_HOST")
			redisPORT := os.Getenv("REDIS_PORT")
			redisURL := fmt.Sprintf("%v:%v", redisHOST, redisPORT)

			ctxRedis := context.Background()

			cliente := redis.NewClient(&redis.Options{
				Addr:     redisURL,
				Password: "", // sin password
				DB:       0,  // usar DB default
			})
			//generando una key random
			randomKey := strings.Replace(uuid.New().String(), "-", "", -1)

			data, err := encoder.Marshal(m)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
			}

			err = cliente.Set(ctxRedis, randomKey, data, 0).Err()
			if err != nil {
				panic(err)
			} else {
				fmt.Println(">> KAFKA: Mensaje guardado en Redis :)")
			}

			/***************************/

		}

	}

}
