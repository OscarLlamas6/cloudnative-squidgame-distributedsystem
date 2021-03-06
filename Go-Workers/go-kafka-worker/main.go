package main

import (
	"context"
	encoder "encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/segmentio/kafka-go"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func isJSON(s string) bool {
	var js map[string]interface{}
	return encoder.Unmarshal([]byte(s), &js) == nil

}

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
		//MinBytes:    0,
		//MaxBytes:    10e6, //10MB
	})

	fmt.Println("Esperando mensajes...")

	for {
		//Consumiento mensajes del Topic
		msg, err := r.ReadMessage(ctxKafKa)
		if err != nil {
			panic("could not read message " + err.Error())
		} else {

			fmt.Println("Kafka >> Mensaje recibido :D")

			if isJSON(string(msg.Value)) {
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

				m["service"] = "Kafka"

				/*GUARDAR EN REDIS*/
				redisHOST := os.Getenv("REDIS_HOST")
				redisPORT := os.Getenv("REDIS_PORT")
				redisPASS := os.Getenv("REDIS_PASS")
				redisURL := fmt.Sprintf("%v:%v", redisHOST, redisPORT)

				ctxRedis := context.Background()
				cliente := redis.NewClient(&redis.Options{
					Addr:     redisURL,
					Password: redisPASS, // sin password
					DB:       0,         // usar DB default
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

				todoCollection := c.Database(mongoDB).Collection(mongoCOL)
				r, err6 := todoCollection.InsertOne(ctxInsert, m)
				if err6 != nil {
					log.Fatalf("Error al guardar logs %v", err)
				} else {
					fmt.Println(">> KAFKA: Log guardado en Mongo con el id: ", r.InsertedID)
				}

				c.Disconnect(ctxInsert)
				cancel()
				/*--------------------FIN DE LOG EN MONGO------------------*/
			} else {
				fmt.Println(string(msg.Value))
			}

		}
	}
}
