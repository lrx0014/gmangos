package config

import (
	"github.com/pelletier/go-toml/v2"
	"os"
	"time"
)

var C Conf

type Conf struct {
	Server ServerConf
	Redis  RedisConf
	MySQL  MySQLConf
}

type ServerConf struct {
	Host         string
	Port         string
	LogPath      string
	LogCacheSize int
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

func LoadConfig(data []byte) Conf {
	C = Conf{}
	err := toml.Unmarshal(data, &C)
	if err != nil {
		panic(err)
	}

	C.parse()
	return C
}

func GetConf() Conf {
	return C
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
