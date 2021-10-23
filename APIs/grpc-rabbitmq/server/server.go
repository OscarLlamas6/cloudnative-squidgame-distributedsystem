package main

import (
	"context"
	encoder "encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"strconv"
	"time"
	"unicode"

	"github.com/joho/godotenv"
	"github.com/streadway/amqp"

	"github.com/OscarLlamas6/grpc-helpers/protos/squidgame"
	"google.golang.org/grpc"
)

type server struct {
	squidgame.UnimplementedSquidGameServiceServer
}

type Atributos struct {
	Mensaje    string `json:"mensaje"`
	Gamenumber string `json:"gamenumber"`
	Gamename   string `json:"gamename"`
	Ganador    string `json:"ganador"`
}

/*ALGORITMOS PARA ELEGIR UN GANADOR*/

func Juego1(maximo int64) int64 {
	rand.Seed(time.Now().UnixNano())
	var winner int64
	for i := 0; i < 6; i++ {
		winner = rand.Int63n(maximo) + 1
	}
	return winner
}

func Juego2(maximo int64) int64 {
	rand.Seed(time.Now().UnixNano())
	var winner int64
	for i := 0; i < 12; i++ {
		winner = rand.Int63n(maximo) + 1
	}
	return winner
}

func Juego3(maximo int64) int64 {
	rand.Seed(time.Now().UnixNano())
	var winner int64
	for i := 0; i < 18; i++ {
		winner = rand.Int63n(maximo) + 1
	}
	return winner
}

func Juego4(maximo int64) int64 {
	rand.Seed(time.Now().UnixNano())
	var winner int64
	for i := 0; i < 24; i++ {
		winner = rand.Int63n(maximo) + 1
	}
	return winner
}

func Juego5(maximo int64) int64 {
	rand.Seed(time.Now().UnixNano())
	var winner int64
	for i := 0; i < 30; i++ {
		winner = rand.Int63n(maximo) + 1
	}
	return winner
}

/*FIN DE ALGORITMOS PARA JUEGOS*/

func (s *server) Play(ctx context.Context, req *squidgame.PlayRequest) (*squidgame.PlayResponse, error) {
	var ganador int64
	gamenumber := req.GetGame().GetGamenumber()
	gamename := req.GetGame().GetGamename()
	players := req.GetGame().GetPlayers()
	// rungames := req.GetGame().GetRungames()
	// concurrence := req.GetGame().GetConcurrence()
	// timeout := req.GetGame().GetTimeout()
	fmt.Println(">> SERVER: Ejecutando juego: ", gamename)
	if isInt(gamenumber) {
		juego, _ := strconv.ParseInt(gamenumber, 10, 64)
		switch juego {
		case 1:
			ganador = Juego1(players)
		case 2:
			ganador = Juego2(players)
		case 3:
			ganador = Juego3(players)
		case 4:
			ganador = Juego4(players)
		case 5:
			ganador = Juego5(players)
		default:
			ganador = Juego5(players)
		}
	} else {
		rand.Seed(time.Now().UnixNano())
		juego := rand.Int63n(5) + 1
		switch juego {
		case 1:
			ganador = Juego1(players)
		case 2:
			ganador = Juego2(players)
		case 3:
			ganador = Juego3(players)
		case 4:
			ganador = Juego4(players)
		case 5:
			ganador = Juego5(players)
		default:
			ganador = Juego1(players)
		}
	}

	fmt.Println(">> SERVER: Juego finalizado!")

	/*ENVIAR A RABBIT*/

	RABBITMQ_HOST := os.Getenv("RABBITMQ_HOST")
	RABBITMQ_PORT := os.Getenv("RABBITMQ_PORT")
	RABBITMQ_USER := os.Getenv("RABBITMQ_USER")
	RABBITMQ_PASS := os.Getenv("RABBITMQ_PASS")
	RABBITMQ_QUEUE := os.Getenv("RABBITMQ_QUEUE")
	RABBITMQ_URL := fmt.Sprintf("amqp://%v:%v@%v:%v/", RABBITMQ_USER, RABBITMQ_PASS, RABBITMQ_HOST, RABBITMQ_PORT)

	conn, errmq := amqp.Dial(RABBITMQ_URL)
	if errmq != nil {
		fmt.Printf("Error: %v\n", errmq)
	}

	defer conn.Close()

	ch, errmq := conn.Channel()
	if errmq != nil {
		fmt.Printf("Error: %v\n", errmq)
	}
	defer ch.Close()

	// We create a Queue to send the message to.
	q, errmq := ch.QueueDeclare(
		RABBITMQ_QUEUE, // name
		false,          // durable
		false,          // delete when unused
		false,          // exclusive
		false,          // no-wait
		nil,            // arguments
	)

	if errmq != nil {
		fmt.Printf("Error: %v\n", errmq)
	}

	mensajeRabbit := "Resultados del juego: " + gamename + " :D"

	// We set the payload for the message
	body := Atributos{
		Mensaje:    mensajeRabbit,
		Gamenumber: "11",
		Gamename:   "Game1121",
		Ganador:    "612121",
	}

	jsonObj, err := encoder.Marshal(body)
	if err != nil {
		fmt.Printf("Error: %v\n", errmq)
	}

	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(jsonObj),
		})

	if err != nil {
		fmt.Printf("Error: %v\n", errmq)
	} else {
		fmt.Println(">> SERVER: Mensaje enviado a RabbitMQ correctamente :D")

	}

	/*************************************************/

	result := "El ganador del juego " + gamename + " es: " + strconv.Itoa(int(ganador))
	res := &squidgame.PlayResponse{
		Message: result,
	}
	fmt.Println(">> SERVER: ", result)
	return res, nil
}

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	grpc_server_host := os.Getenv("RABBIT_SERVER_HOST")
	instance_name := os.Getenv("RABBIT_SERVER_NAME")
	fmt.Println(">> -------- ", instance_name, " --------")
	fmt.Println(">> SERVER: Iniciando servidor http en ", grpc_server_host)
	list, err := net.Listen("tcp", grpc_server_host)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	squidgame.RegisterSquidGameServiceServer(s, &server{})

	if err := s.Serve(list); err != nil {
		log.Fatalf("faile to serve: %v", err)
	}
}

func isInt(s string) bool {
	for _, c := range s {
		if !unicode.IsDigit(c) {
			return false
		}
	}
	return true
}
