package main

import (
	"encoding/json"
	"fmt"
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
)

type AMQPManager struct {
	conn    *amqp.Connection
	channel *amqp.Channel
}

// NewAMQPManager creates and returns a new AMQP manager instance
func NewAMQPManager() (*AMQPManager, error) {
	// Get credentials from environment
	amqpURL := os.Getenv("AMQP_URL")
	if amqpURL == "" {
		return nil, fmt.Errorf("AMQP_URL not set in environment")
	}

	conn, err := amqp.Dial(amqpURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to RabbitMQ: %v", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("failed to open channel: %v", err)
	}

	return &AMQPManager{
		conn:    conn,
		channel: ch,
	}, nil
}

// Close closes the AMQP connection and channel
func (m *AMQPManager) Close() {
	if m.channel != nil {
		m.channel.Close()
	}
	if m.conn != nil {
		m.conn.Close()
	}
}

// DeclareQueue declares a queue
func (m *AMQPManager) DeclareQueue(queueName string) error {
	// Declare the queue
	_, err := m.channel.QueueDeclare(
		queueName, // name
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		true,      // no-wait
		nil,       // arguments
	)
	if err != nil {
		return fmt.Errorf("failed to declare queue: %v", err)
	}

	return nil
}

// QueueMessage queues a message to the specified queue
func (m *AMQPManager) QueueMessage(queueName string, data interface{}) error {
	// Convert data to JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("failed to marshal data: %v", err)
	}

	// Publish the message
	err = m.channel.Publish(
		"",        // exchange
		queueName, // routing key
		false,     // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        jsonData,
		})
	if err != nil {
		return fmt.Errorf("failed to publish message: %v", err)
	}
	return nil
}
