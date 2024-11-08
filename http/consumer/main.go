package main

import (
	"fmt"
	"io"
	"log"
	"math/rand/v2"
	"net/http"
	"os"
	"strings"
	"time"
)

func messageHandler(w http.ResponseWriter, r *http.Request) {
	buf := new(strings.Builder)
	_, err := io.Copy(buf, r.Body)
	if err != nil {
		fmt.Printf("Error in reading req - %v", err)
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	message := buf.String()
	fmt.Println("Got Message - ", message)

	time.Sleep(2 * time.Second)

	defer r.Body.Close()

	randomNumber := rand.IntN(3)
	if randomNumber == 2 {
		log.Printf("RETURN - ERROR NO RETRY - %v ------", message)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		// https://docs.dapr.io/reference/api/pubsub_api/#provide-routes-for-dapr-to-deliver-topic-events
		writeResponse(w, `{"status": "DROP"}`)
		return
	}
	if randomNumber == 1 {
		log.Printf("RETURN - RETRIABLE ERROR - %v ------", message)

		w.Header().Set("Content-Type", "application/json")
		// Non 2XX is retried. Also can set the staus as 200 and give the response as {"status": "RETRY"} or {"status": "ANY-STRING"} (refer below doc)
		w.WriteHeader(http.StatusInternalServerError)
		// https://docs.dapr.io/reference/api/pubsub_api/#provide-routes-for-dapr-to-deliver-topic-events
		writeResponse(w, `{"status": "RETRY"}`)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	writeResponse(w, `{"success": "true"}`)
}

func writeResponse(w http.ResponseWriter, responseString string) {
	response := []byte("Hello, world!")
	_, err := w.Write(response)
	if err != nil {
		fmt.Println("Failed to write response - ", responseString)
		fmt.Printf("Error - %v", err)
	}
}

func main() {
	port := "6002"
	if len(os.Args) > 1 {
		port = os.Args[1]
	}

	// Route for handling messages
	http.HandleFunc("/events", messageHandler)

	fmt.Println("Subscriber listening on port", port)
	srv := &http.Server{Addr: fmt.Sprintf(":%s", port)}
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
