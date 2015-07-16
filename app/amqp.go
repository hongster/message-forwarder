// Provide AMQP specific information from config file.
package app

import (
	"fmt"
)

// Get exhange name from config file.
func ExchangeName() string {
	return Config.StringDefault("message", "exchange", "callback")
}

// Generate AMQP URL based on configurations.
// TODO Support SSL connection
func AMQPURL() string {
	return fmt.Sprintf("amqp://%s:%s@%s:%s/%s",
		Config.StringDefault("DEFAULT", "amqp_username", ""),
		Config.StringDefault("DEFAULT", "amqp_password", ""),
		Config.StringDefault("DEFAULT", "amqp_host", ""),
		Config.StringDefault("DEFAULT", "amqp_port", ""),
		Config.StringDefault("message", "vhost", ""))
}
