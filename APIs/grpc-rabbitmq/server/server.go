package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"

	"github.com/OscarLlamas6/grpc-helpers/protos/squidgame"
	"google.golang.org/grpc"
)

type server struct {
	squidgame.UnimplementedSquidGameServiceServer
}

func (s *server) Play(ctx context.Context, req *squidgame.PlayRequest) (*squidgame.PlayResponse, error) {
	fmt.Println(">> SERVER: Ejecutando juego :o")
	// gamenumber := req.GetGame().GetGamenumber()
	gamename := req.GetGame().GetGamename()
	// players := req.GetGame().GetPlayers()
	// rungames := req.GetGame().GetRungames()
	// concurrence := req.GetGame().GetConcurrence()
	// timeout := req.GetGame().GetTimeout()

	/*IMPORTANTEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEE
	GENERAR GANADOR Y ENVIAR A COLA
	*/

	fmt.Println(">> SERVER: Juego finalizado!")

	result := "RABBIT gRPC Server >> El ganador del juego " + gamename + " es: X"
	res := &squidgame.PlayResponse{
		Message: result,
	}
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
