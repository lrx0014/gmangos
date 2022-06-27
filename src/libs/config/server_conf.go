package config

import "fmt"

const (
	DefaultAddr = "127.0.0.1"
	DefaultPort = "8888"
)

func (c *ServerConf) ParseConfig() (err error) {
	if c.Addr == "" {
		c.Addr = DefaultAddr
	}
	if c.Port == "" {
		c.Port = DefaultPort
	}

	return
}

func (c *ServerConf) Address() string {
	return fmt.Sprintf("%s:%s", c.Addr, c.Port)
}
