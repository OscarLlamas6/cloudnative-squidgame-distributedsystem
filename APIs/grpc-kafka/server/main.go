package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
	"unicode"

	encoder "encoding/json"

	"github.com/OscarLlamas6/grpc-helpers/protos/squidgame"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/segmentio/kafka-go"
	"google.golang.org/grpc"
)

type server struct {
	squidgame.UnimplementedSquidGameServiceServer
}

type Atributos struct {
	Gamenumber string `json:"gamenumber"`
	Gamename   string `json:"gamename"`
	Ganador    string `json:"ganador"`
	Players    string `json:"players"`
	Request    string `json:"request"`
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
	request := req.GetGame().GetRequest()
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

	/*ENVIAR A KAFKA*/

	ctxKafKa := context.Background()

	kafkaHOST := os.Getenv("KAFKA_HOST")
	kafkaPORT := os.Getenv("KAFKA_PORT")
	kafkaTOPIC := os.Getenv("KAFKA_TOPIC")
	brokerURL := fmt.Sprintf("%v:%v", kafkaHOST, kafkaPORT)

	l := log.New(os.Stdout, ">> KAFKA: ", 0)
	// intialize the writer with the broker addresses, and the topic
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{brokerURL},
		Topic:   kafkaTOPIC,
		// assign the logger to the writer
		Logger: l,
	})

	// We set the payload for the message
	body := Atributos{
		Gamenumber: gamenumber,
		Gamename:   gamename,
		Ganador:    strconv.Itoa(int(ganador)),
		Players:    strconv.Itoa(int(players)),
		Request:    strconv.Itoa(int(request)),
	}

	jsonObj, err := encoder.Marshal(body)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	randomKey := strings.Replace(uuid.New().String(), "-", "", -1)

	err = w.WriteMessages(ctxKafKa, kafka.Message{
		Key: []byte(randomKey),
		// create an arbitrary message payload for the value
		Value: []byte(string(jsonObj)),
	})
	if err != nil {
		panic("Error al mandar mensaje a Kafka: " + err.Error())
	} else {
		fmt.Println(">> SERVER: Mensaje enviado a KAFKA correctamente :D")
	}

	/*********************/

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

	grpc_server_port := os.Getenv("KAFKA_SERVER_PORT")
	grpc_server_host := fmt.Sprintf(":%v", grpc_server_port)
	instance_name := os.Getenv("KAFKA_SERVER_NAME")
	fmt.Println(">> -------- ", instance_name, " --------")
	fmt.Println(">> SERVER: Iniciando servidor gRPC en ", grpc_server_host)
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
