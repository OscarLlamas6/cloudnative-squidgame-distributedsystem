package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"cloud.google.com/go/pubsub"
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

	sub := client.Subscription(subID)
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	cm := make(chan *pubsub.Message)
	defer close(cm)

	go func() {
		for msg := range cm {

			/*GUARDAR INFO EN REDIS*/

			fmt.Printf("Got message :%q\n", string(msg.Data))
			fmt.Println("Attributes:")
			for key, value := range msg.Attributes {
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
