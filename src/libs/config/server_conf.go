package config

import (
	"fmt"
	"time"
)

const (
	DefaultServerHost    = "127.0.0.1"
	DefaultServerPort    = "8888"
	DefaultServerLogPath = "./logs/%v.log"
)

func (c *ServerConf) ParseConfig() (err error) {
	if c.Host == "" {
		c.Host = DefaultServerHost
	}
	if c.Port == "" {
		c.Port = DefaultServerPort
	}
	if c.LogPath == "" {
		c.LogPath = fmt.Sprintf(DefaultServerLogPath, time.Now().Unix())
	}
	if c.LogCacheSize == 0 {
		c.LogCacheSize = 1024
	}

	return
}

func (c *ServerConf) Address() string {
	return fmt.Sprintf("%s:%s", c.Host, c.Port)
}
