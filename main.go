package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// queue name
	queueName := os.Getenv("QUEUE_NAME")
	if queueName == "" {
		log.Fatalf("QUEUE_NAME not set in environment")
	}

	am, err := NewAMQPManager()
	if err != nil {
		log.Fatalf("Failed to create AMQP manager: %v", err)
	} else {
		log.Printf("AMQP manager created")
	}
	defer am.Close()

	err = am.DeclareQueue(queueName)
	if err != nil {
		log.Fatalf("Failed to declare queue: %v", err)
	} else {
		log.Printf("Queue %s declared", queueName)
	}

	s := NewServer(am, queueName)
	s.Logger.Fatal(s.Start(":8080"))
}
