package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"github.com/OscarLlamas6/grpc-helpers/protos/squidgame"
	"google.golang.org/grpc"
)

// SingleGame structure
type SingleGame struct {
	Gamenumber  string `json:"gamenumber"`
	Gamename    string `json:"gamename"`
	Players     int64  `json:"players"`
	Rungames    int64  `json:"rungames"`
	Concurrence int64  `json:"concurrence"`
	Timeout     int64  `json:"timeout"`
}

func http_server(w http.ResponseWriter, r *http.Request) {
	instance_name := os.Getenv("PUBSUB_CLIENT_NAME")
	fmt.Println(">> CLIENT: Manejando peticion HTTP CLIENTE: ", instance_name)
	// Comprobamos que el path sea exactamente '/' sin parámetros
	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	// Comprobamos el tipo de peticion HTTP
	switch r.Method {

	case "POST":
		fmt.Println(">> CLIENT: Iniciando envio de mensajes")
		// Si existe un error con la forma enviada entonces no seguir
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}

		var g SingleGame
		err := json.NewDecoder(r.Body).Decode(&g)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		grpc_server_host := os.Getenv("PUBSUB_SERVER_HOST")
		fmt.Println(">> CLIENT: Iniciando cliente gRPC")
		fmt.Println(">> CLIENT: Iniciando conexion con el servidor gRPC ", grpc_server_host)

		conn, err := grpc.Dial(grpc_server_host, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("could not connect: %v", err)
		}
		defer conn.Close()
		c := squidgame.NewSquidGameServiceClient(conn)
		sendGame(c, g.Gamenumber, g.Gamename, g.Players, g.Rungames, g.Concurrence, g.Timeout)
		fmt.Fprintf(w, "¡Juego enviado al SQUID GAME gRPC Server!\n")

	default:
		fmt.Fprintf(w, "Metodo %s no soportado \n", r.Method)
		return
	}
}

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	instance_name := os.Getenv("PUBSUB_CLIENT_NAME")
	client_host := os.Getenv("PUBSUB_CLIENT_HOST")

	fmt.Println(">> --------  ", instance_name, " --------")
	fmt.Println(">> CLIENT: Iniciando servidor http en ", client_host)
	http.HandleFunc("/", http_server)
	if err := http.ListenAndServe(client_host, nil); err != nil {
		log.Fatal(err)
	}

}

func sendGame(c squidgame.SquidGameServiceClient, number string, name string, players int64, rungames int64, concurrence int64, timeout int64) {
	fmt.Println(">> CLIENT: Enviando Squid Game a gRPC Server ")
	req := &squidgame.PlayRequest{
		Game: &squidgame.Game{
			Gamenumber:  number,
			Gamename:    name,
			Players:     players,
			Rungames:    rungames,
			Concurrence: concurrence,
			Timeout:     timeout,
		},
	}

	res, err := c.Play(context.Background(), req)

	if err != nil {
		log.Fatalf("Se ha producido un error: %v :(", err)
	}

	fmt.Printf(">> CLIENT: %v\n", res.Message)
}