package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"math/rand/v2"
	"net/http"
	"os"
	"time"

	"github.com/dinumathai/dapr-pubsub/appconst"

	"github.com/dapr/go-sdk/service/common"
	daprd "github.com/dapr/go-sdk/service/http"
)

// code
var sub = &common.Subscription{
	PubsubName: appconst.PUBSUB_NAME,
	Topic:      appconst.TOPIC_NAME,
	Route:      "/events", // The Route of the application
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

func eventHandler(ctx context.Context, e *common.TopicEvent) (bool, error) {
	log.Printf("Subscriber received: %v", e.Data)
	time.Sleep(2 * time.Second)
	randomNumber := rand.IntN(3)
	if randomNumber == 2 {
		log.Printf("RETURN - ERROR NO RETRY - %v ------", e.Data)
		return false, errors.New("pls DO NOT do a retry")
	}
	if randomNumber == 1 {
		log.Printf("RETURN - RETRIABLE ERROR - %v ------", e.Data)
		return true, errors.New("pls do a retry")
	}
	return false, nil
}
