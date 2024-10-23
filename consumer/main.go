package main

import (
	"context"
	"dapr-pubsub/appconst"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/dapr/go-sdk/service/common"
	daprd "github.com/dapr/go-sdk/service/http"
)

// code
var sub = &common.Subscription{
	PubsubName: appconst.PUBSUB_NAME,
	Topic:      appconst.TOPIC_NAME,
	Route:      "/checkout",
}

func main() {
	port := "6002"
	if len(os.Args) > 1 {
		port = os.Args[1]
	}
	flag.Parse()
	listeningPort := fmt.Sprintf(":%s", port)
	log.Printf("Starting Server at %s", listeningPort)
	s := daprd.NewService(listeningPort)
	//Subscribe to a topic
	if err := s.AddTopicEventHandler(sub, eventHandler); err != nil {
		log.Fatalf("error adding topic subscription: %v", err)
	}
	if err := s.Start(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("error listening: %v", err)
	}
}

func eventHandler(ctx context.Context, e *common.TopicEvent) (retry bool, err error) {
	log.Printf("Subscriber received: %v", e.Data)
	time.Sleep(3 * time.Second)
	log.Printf("------------------------")
	return false, nil
}
