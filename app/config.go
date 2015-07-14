// Provide access to configurations in /etc/message-forwarder.cfg
package app

import (
	"github.com/hongster/message-forwarder/config"
	"fmt"
)

// Ini-style config file, /etc/message-forwarder.cfg
const CONFIG_FILE = "/etc/message-forwarder.cfg"

// Reference to config reader
var Config *config.Reader

func init() {
	var err error
	Config, err = config.NewReader(CONFIG_FILE)
	if err != nil {
		panic(fmt.Sprintf("Unable to access config file: %v", err))
	}
}
