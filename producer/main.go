package main

//dependencies
import (
	"context"
	"dapr-pubsub/appconst"
	"log"
	"math/rand"
	"strconv"
	"time"

	dapr "github.com/dapr/go-sdk/client"
)

func main() {
	for i := 0; i < 10; i++ {
		time.Sleep(5000)
		orderId := rand.Intn(1000-1) + 1
		client, err := dapr.NewClient()
		if err != nil {
			panic(err)
		}
		defer client.Close()
		ctx := context.Background()
		//Using Dapr SDK to publish a topic
		if err := client.PublishEvent(ctx, appconst.PUBSUB_NAME, appconst.TOPIC_NAME, []byte(strconv.Itoa(orderId))); err != nil {
			panic(err)
		}

		log.Println("Published data: " + strconv.Itoa(orderId))
	}
}
