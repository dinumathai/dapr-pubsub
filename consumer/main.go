package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Order struct {
	OrderID string `json:"orderId"`
	Amount  int    `json:"amount"`
}

func main() {
	http.HandleFunc("/orders", func(w http.ResponseWriter, r *http.Request) {
		var order Order
		if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}
		fmt.Printf("Received order: %+v\n", order)
	})

	fmt.Println("Starting consumer, waiting for messages...")
	if err := http.ListenAndServe(":6000", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
