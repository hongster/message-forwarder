package main

import (
	"github.com/hongster/message-forwarder/app"
	"github.com/hongster/message-forwarder/worker"
	"github.com/streadway/amqp"
)

const CONSUMER_ID = "message-forwarder"

func main() {
	// Setup connection
	app.Logger.Info("Connecting to AMQP server...")
	connection, err := amqp.Dial(app.AMQPURL())
	if err != nil {
		app.Logger.Error("Can't connect to AMQP server: %v", err)
		return
	}
	app.Logger.Info("Connected")

	defer func() {
		connection.Close()
		app.Logger.Info("Connection closed")
	}()

	// Setup channel
	channel, err := connection.Channel()
	if err != nil {
		app.Logger.Error("Can't get channel: %v", err)
		return
	}

	// Graceful cancelling and closing channel
	defer func() {
		channel.Cancel(CONSUMER_ID, false)
		app.Logger.Info("Connection cancelled")
	}()

	// Request message to be ACK upon consumption
	deliveryChan, err := channel.Consume(app.ExchangeName(), CONSUMER_ID, false, false, false, false, nil)
	if err != nil {
		app.Logger.Error("Can't consume: %v", err)
		return
	}

	// Process task messages
	for delivery := range deliveryChan {

		// Each message is processed in a Go routine. ACK will be sent if upon
		// successful processing, NACK otherwise.
		go func() {
			err = worker.Process(delivery)
			if err != nil {
				app.Logger.Error("%s", err)
				delivery.Nack(false, false)
				return
			}

			delivery.Ack(false)
		}()

	}
}
