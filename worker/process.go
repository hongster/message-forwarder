// Forward message body to callback-URL defined in message header.
package worker

import (
	"github.com/hongster/message-forwarder/app"
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

	if delivery.UserId == "" {
		return fmt.Errorf("Dropping message due to missing user ID")
	}

	app.Logger.Info("Forwarding message from %s to %s", delivery.UserId, callbackURL)
	resp, err := http.Post(callbackURL.(string), delivery.ContentType, bytes.NewBuffer(delivery.Body))
	if err != nil {
		return err
	}

	app.Logger.Info("Response status code: %d", resp.StatusCode)
	resp.Body.Close()
	return nil
}
