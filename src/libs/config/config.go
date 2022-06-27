package config

import (
	"github.com/pelletier/go-toml/v2"
	"os"
	"time"
)

type Conf struct {
	Server ServerConf
	Redis  RedisConf
	MySQL  MySQLConf
}

type ServerConf struct {
	Host string
	Port string
}

type RedisConf struct {
	Host            string
	Port            string
	MaxIdle         int
	MaxActive       int
	IdleTimeout     time.Duration
	MaxConnLifetime time.Duration
}

type MySQLConf struct {
	DSN         string
	Active      int
	Idle        int
	IdleTimeout time.Duration
}

func ReadFile(path string) (data []byte) {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	return
}

func LoadConfig(data []byte) *Conf {
	c := &Conf{}
	err := toml.Unmarshal(data, c)
	if err != nil {
		panic(err)
	}

	c.parse()
	return c
}

func (c *Conf) parse() {
	err := c.Server.ParseConfig()
	if err != nil {
		panic(err)
	}

	err = c.Redis.ParseConfig()
	if err != nil {
		panic(err)
	}

	err = c.MySQL.ParseConfig()
	if err != nil {
		panic(err)
	}
}
