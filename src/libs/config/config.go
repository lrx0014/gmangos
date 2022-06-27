package config

import (
	"github.com/pelletier/go-toml/v2"
	"os"
)

type Conf struct {
	Server *ServerConf
	Redis  *RedisConf
	DB     *DBConf
}

type ServerConf struct {
	Addr string
	Port string
}

type RedisConf struct {
}

type DBConf struct {
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

	return c
}
