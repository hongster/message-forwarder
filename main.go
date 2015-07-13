package main

import (
	"github.com/hongster/message-forwarder/config"
	"github.com/hongster/message-forwarder/logger"
	"github.com/hongster/message-forwarder/worker"
	"github.com/streadway/amqp"
	"fmt"
)

const CONSUMER_ID = "message-forwarder"

func main() {
	// Setup connection
	connection, err := amqp.Dial(amqpURL())
	if err != nil {
		logger.Error("Can't connect to AMQP server: %v", err)
		return
	}

	defer func() {
		connection.Close()
		logger.Info("Connection closed")
	}()

	// Setup channel
	channel, err := connection.Channel()
	if err != nil {
		logger.Error("Can't get channel: %v", err)
		return
	}

	// Graceful cancelling and closing channel
	defer func() {
		channel.Cancel(CONSUMER_ID, false)
		logger.Info("Connection cancelled")
	}()

	// Request message to be ACK upon consumption
	deliveryChan, err := channel.Consume(exchangeName(), CONSUMER_ID, false, false, false, false, nil)
	if err != nil {
		logger.Error("Can't consume: %v", err)
		return
	}

	// Process task messages
	for delivery := range deliveryChan {

		// Each message is processed in a Go routine. ACK will be sent if upon
		// successful processing, NACK otherwise.
		go func() {
			err = worker.Process(delivery)
			if err != nil {
				logger.Error("%s", err)
				delivery.Nack(false, false)
				return
			}

			delivery.Ack(false)
		}()

	}
}

// Get exhange name from config file.
func exchangeName() string {
	configReader := config.NewReader()
	return configReader.StringDefault("message", "exchange", "callback")
}

// Generate AMQP URL based on configurations.
// TODO Support SSL connection
func amqpURL() string {
	configReader := config.NewReader()

	return fmt.Sprintf("amqp://%s:%s@%s:%s/%s",
		configReader.StringDefault("DEFAULT", "amqp_username", ""),
		configReader.StringDefault("DEFAULT", "amqp_password", ""),
		configReader.StringDefault("DEFAULT", "amqp_host", ""),
		configReader.StringDefault("DEFAULT", "amqp_port", ""),
		configReader.StringDefault("message", "vhost", ""))
}
