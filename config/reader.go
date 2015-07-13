/*
Read configuration INI config from /etc/message-forwarder.cfg
*/
package config

import (
	"fmt"
	"github.com/robfig/config"
)

const CONFIG_FILE = "/etc/message-forwarder.cfg"

type Reader struct {
	config *config.Config
}

func NewReader() *Reader {
	config, err := config.ReadDefault(CONFIG_FILE)
	if err != nil {
		config = nil
	}

	return &Reader{
		config: config,
	}
}

func (reader *Reader) Bool(section string, option string) (value bool, err error) {
	if reader.config == nil {
		return false, fmt.Errorf("Missing config file: %s", CONFIG_FILE)
	}

	return reader.config.Bool(section, option)
}

/*
Return default value if config option no found.
*/
func (reader *Reader) BoolDefault(section string, option string, def bool) (value bool) {
	value, err := reader.Bool(section, option)
	if err != nil {
		return def
	}

	return value
}

func (reader *Reader) Int(section string, option string) (value int, err error) {
	if reader.config == nil {
		return 0, fmt.Errorf("Missing config file: %s", CONFIG_FILE)
	}

	return reader.config.Int(section, option)
}

/*
Return default value if config option no found.
*/
func (reader *Reader) IntDefault(section string, option string, def int) (value int) {
	value, err := reader.Int(section, option)
	if err != nil {
		return def
	}

	return value
}

func (reader *Reader) String(section string, option string) (value string, err error) {
	if reader.config == nil {
		return "", fmt.Errorf("Missing config file: %s", CONFIG_FILE)
	}

	return reader.config.String(section, option)
}

/*
Return default value if config option no found.
*/
func (reader *Reader) StringDefault(section string, option string, def string) (value string) {
	value, err := reader.String(section, option)
	if err != nil {
		return def
	}

	return value
}

func (reader *Reader) Float(section string, option string) (value float64, err error) {
	if reader.config == nil {
		return 0, fmt.Errorf("Missing config file: %s", CONFIG_FILE)
	}

	return reader.config.Float(section, option)
}

/*
Return default value if config option no found.
*/
func (reader *Reader) FloatDefault(section string, option string, def float64) (value float64) {
	value, err := reader.Float(section, option)
	if err != nil {
		return def
	}

	return value
}
