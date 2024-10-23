package main

//dependencies
import (
	"context"
	"dapr-pubsub/appconst"
	"log"
	"time"

	dapr "github.com/dapr/go-sdk/client"
)

func main() {
	for i := 0; i < 10; i++ {
		time.Sleep(2 * time.Second)
		client, err := dapr.NewClient()
		if err != nil {
			panic(err)
		}
		defer client.Close()
		ctx := context.Background()
		orderId := time.Now().Format(time.RFC3339)
		//Using Dapr SDK to publish a topic
		if err := client.PublishEvent(ctx, appconst.PUBSUB_NAME, appconst.TOPIC_NAME, []byte(orderId)); err != nil {
			panic(err)
		}

		log.Println("Published data: " + orderId)
	}
}
