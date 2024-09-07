package main

import (
	"fmt"
	"log"
)

func main() {
	server := NewAPIServer(":3000")

	store, err := NewPostgresStore()
	if err != nil {
		log.Fatalf("Error creating store: %v", err)
	}

	if err := store.Init(); err != nil {
		log.Fatalf("Error initializing store: %v", err)
	}

	fmt.Printf("%+v\n", store)

	server.Run()
}
