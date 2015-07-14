package config

import (
	"github.com/robfig/config"
)

type Reader struct {
	config *config.Config
}

// filepath E.g. ""/etc/someconfig.cfg"
func NewReader(filepath string) (reader *Reader, err error) {
	config, err := config.ReadDefault(filepath)
	if err != nil {
		return nil, err
	}

	return &Reader{config: config}, nil
}

func (reader *Reader) Bool(section string, option string) (value bool, err error) {
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
