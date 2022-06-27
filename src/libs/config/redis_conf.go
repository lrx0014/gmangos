package config

import "fmt"

const (
	DefaultRedisHost = "127.0.0.1"
	DefaultRedisPort = "6379"
)

func (c *RedisConf) ParseConfig() (err error) {
	if c.Host == "" {
		c.Host = DefaultRedisHost
	}
	if c.Port == "" {
		c.Port = DefaultRedisPort
	}

	return
}

func (c *RedisConf) Address() string {
	return fmt.Sprintf("%s:%s", c.Host, c.Port)
}
