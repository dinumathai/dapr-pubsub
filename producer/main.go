package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

const pubsubName = "my-pubsub"
const topicName = "orders"

type Order struct {
	OrderID string `json:"orderId"`
	Amount  int    `json:"amount"`
}

func main() {
	order := Order{
		OrderID: "123",
		Amount:  250,
	}

	data, err := json.Marshal(order)
	if err != nil {
		fmt.Println("Error marshalling order:", err)
		os.Exit(1)
	}

	url := fmt.Sprintf("http://localhost:3500/v1.0/publish/%s/%s", pubsubName, topicName)

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(data))
	if err != nil {
		fmt.Println("Error publishing message:", err)
		os.Exit(1)
	}

	defer resp.Body.Close()
	fmt.Println("Published message to topic:", topicName)
}
