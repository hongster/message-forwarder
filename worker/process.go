// Forward message body to callback-URL defined in message header.
package worker

import (
	"github.com/hongster/message-forwarder/logger"
	"github.com/streadway/amqp"
	"bytes"
	"fmt"
	"net/http"
)

// Required in message header
const CALLBACK_URL_HEADER = "callback-url"

// **Expecting** "callback-url" in message header. It defines the URL for
// forwarding message content to.
// It is *advisable* to define message content-type.
func Process(delivery amqp.Delivery) (err error) {
	callbackURL, ok := delivery.Headers[CALLBACK_URL_HEADER]
	if !ok {
		return fmt.Errorf("Missing %s message header", CALLBACK_URL_HEADER)
	}

	logger.Info("Forwarding message to %s", callbackURL)
	_, err = http.Post(callbackURL.(string), delivery.ContentType, bytes.NewBuffer(delivery.Body))
	if err != nil {
		return err
	}

	return nil
}
