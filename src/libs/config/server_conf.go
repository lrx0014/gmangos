package config

import "fmt"

const (
	DefaultServerHost = "127.0.0.1"
	DefaultServerPort = "8888"
)

func (c *ServerConf) ParseConfig() (err error) {
	if c.Host == "" {
		c.Host = DefaultServerHost
	}
	if c.Port == "" {
		c.Port = DefaultServerPort
	}

	return
}

func (c *ServerConf) Address() string {
	return fmt.Sprintf("%s:%s", c.Host, c.Port)
}
