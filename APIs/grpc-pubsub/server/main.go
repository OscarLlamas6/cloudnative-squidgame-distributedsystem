package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"strconv"
	"time"
	"unicode"

	"cloud.google.com/go/pubsub"
	"github.com/joho/godotenv"

	"github.com/OscarLlamas6/grpc-helpers/protos/squidgame"
	"google.golang.org/grpc"
)

type server struct {
	squidgame.UnimplementedSquidGameServiceServer
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

	/*ENVIAR A PUBSUB*/

	credentials_path := os.Getenv("PUBSUB_KEY_PATH")
	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", credentials_path)
	projectID := os.Getenv("PUBSUB_PROJECT")
	topicID := os.Getenv("GOLANG_TOPIC")

	ctxPS := context.Background()
	client, error2 := pubsub.NewClient(ctxPS, projectID)
	if error2 != nil {
		log.Fatal("Error creando cliente")
	}
	defer client.Close()

	t := client.Topic(topicID)

	mensajePS := "Resultados del juego: " + gamename + " :D"

	resultPS := t.Publish(ctxPS, &pubsub.Message{
		Data: []byte(mensajePS),
		Attributes: map[string]string{
			"gamenumber": gamenumber,
			"gamename":   gamename,
			"ganador":    strconv.Itoa(int(ganador)),
			"players":    strconv.Itoa(int(players)),
			"request":    strconv.Itoa(int(request)),
		},
	})

	id, error3 := resultPS.Get(ctxPS)
	if error3 != nil {
		fmt.Printf("Error: %v", error3)
	} else {
		fmt.Printf(">> SERVER: Mensaje enviado con el id: %v\n", id)

	}

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

	grpc_server_port := os.Getenv("PUBSUB_SERVER_PORT")
	grpc_server_host := fmt.Sprintf(":%v", grpc_server_port)
	instance_name := os.Getenv("PUBSUB_SERVER_NAME")
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
