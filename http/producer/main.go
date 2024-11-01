package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/dinumathai/dapr-pubsub/appconst"
)

func publishMessage(message string, daprPort string) error {
	url := appconst.MakeDAPRSidecarPubURL(appconst.PUBSUB_NAME, appconst.TOPIC_NAME, daprPort, true)

	payloadBytes := []byte(fmt.Sprintf(`{"the-data": "%s"}`, message))
	// Send the HTTP POST request to Dapr publish endpoint
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(payloadBytes))
	if err != nil {
		return fmt.Errorf("failed to publish message: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("unexpected response from Dapr: %s", resp.Status)
	}

	fmt.Println("Message published:", string(payloadBytes))
	return nil
}

func main() {
	port := "3500" // Default port for Kubernetes sidecar
	if len(os.Args) > 1 {
		port = os.Args[1]
	}
	for {
		time.Sleep(1 * time.Second)
		message := fmt.Sprintf("Hello - %s", time.Now().Format(time.RFC3339))
		if err := publishMessage(message, port); err != nil {
			log.Fatalf("Error publishing message: %v", err)
		}
	}
}
