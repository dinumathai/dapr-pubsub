package main

//dependencies
import (
	"context"
	"log"
	"os"
	"time"

	"github.com/dinumathai/dapr-pubsub/appconst"

	dapr "github.com/dapr/go-sdk/client"
)

func main() {
	daprGrpcPort := "50001" // Default port for Kubernetes sidecar
	if len(os.Args) > 1 {
		daprGrpcPort = os.Args[1]
	}

	for {
		time.Sleep(1 * time.Second)
		publishData(daprGrpcPort)
	}
}

func publishData(daprGrpcPort string) {
	client, err := dapr.NewClientWithPort(daprGrpcPort)
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
