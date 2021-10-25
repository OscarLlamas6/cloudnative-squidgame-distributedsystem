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
	"github.com/streadway/amqp"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	//Leo variables de entorno
	RABBITMQ_HOST := os.Getenv("RABBITMQ_HOST")
	RABBITMQ_PORT := os.Getenv("RABBITMQ_PORT")
	RABBITMQ_USER := os.Getenv("RABBITMQ_USER")
	RABBITMQ_PASS := os.Getenv("RABBITMQ_PASS")
	RABBITMQ_QUEUE := os.Getenv("RABBITMQ_QUEUE")
	RABBITMQ_URL := fmt.Sprintf("amqp://%v:%v@%v:%v/", RABBITMQ_USER, RABBITMQ_PASS, RABBITMQ_HOST, RABBITMQ_PORT)
	//Creo la conexion a RabbitMQ
	conn, err := amqp.Dial(RABBITMQ_URL)
	if err != nil {
		log.Fatal("Error creando conexion")
	}
	//Dejo listo el cierre de la conexion
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal("Error creando channel")
	}
	defer ch.Close()

	//declaro la cola de rabbitMQ a donde mandare los mensajes
	q, err := ch.QueueDeclare(
		RABBITMQ_QUEUE, // name
		false,          // durable
		false,          // delete when unused
		false,          // exclusive
		false,          // no-wait
		nil,            // arguments
	)
	if err != nil {
		log.Fatal("Error al crear la cola.")
	}
	//consumo los mensajes que estan en la cola
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		log.Fatal("Error al obtener mensajes de la cola.")
	}

	forever := make(chan bool)

	go func() {
		//Iterando sobre cada mensaje recbido
		for d := range msgs {
			fmt.Println("RabbitMQ >> Mensaje recibido :D")
			//Creo un map para mappea el arreglo de bytes del cuerpo del mensaje
			m := make(map[string]string)
			err2 := encoder.Unmarshal(d.Body, &m)
			if err2 != nil {
				fmt.Printf("Error: %v\n", err)
			}
			//Itero en cada llave/valor del map
			for key, value := range m {
				fmt.Printf("%s = %s\n", key, value)
			}

			m["service"] = "RabbitMQ"

			/*GUARDAR EN REDIS*/
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

			data, err := encoder.Marshal(m)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
			}

			err = cliente.Set(ctxRedis, randomKey, data, 0).Err()
			if err != nil {
				panic(err)
			} else {
				fmt.Println("RabbitMQ >> Mensaje guardado en Redis :)")
			}

			/*GUARDAR LOGS EN MONGO*/

			mongoHOST := os.Getenv("MONGO_HOST")
			mongoPORT := os.Getenv("MONGO_PORT")
			mongoDB := os.Getenv("MONGO_DB")
			mongoCOL := os.Getenv("MONGO_COL")
			mongoURL := fmt.Sprintf("mongodb://%v:%v", mongoHOST, mongoPORT)

			ctxMongo, cancel := context.WithTimeout(context.Background(), time.Second*10)
			//defer cancel()

			clientOptions := options.Client().ApplyURI(mongoURL).SetDirect(true)

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
			//defer c.Disconnect(ctxInsert)

			todoCollection := c.Database(mongoDB).Collection(mongoCOL)
			r, err := todoCollection.InsertOne(ctxInsert, m)
			if err != nil {
				log.Fatalf("Error al guardar logs %v", err)
			} else {
				fmt.Println("RabbitMQ >> Log guardado en Mongo con el id: ", r.InsertedID)
			}

			cancel()
			/*FIN DE LOG EN MONGO*/
		}
	}()

	log.Printf("Esperando mensajes...")
	<-forever

}
