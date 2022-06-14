package tcp

import "fmt"

const (
	DefaultAddr = "127.0.0.1"
	DefaultPort = "8888"
)

type Conf struct {
	Addr string
	Port string
}

func (c *Conf) parseConfig() (err error) {
	if c.Addr == "" {
		c.Addr = DefaultAddr
	}
	if c.Port == "" {
		c.Port = DefaultPort
	}

	return
}

func (c *Conf) address() string {
	return fmt.Sprintf("%s:%s", c.Addr, c.Port)
}
