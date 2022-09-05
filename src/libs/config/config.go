package config

import (
	"github.com/pelletier/go-toml/v2"
	"gmangos/src/libs/utils"
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
	Host    string
	Port    string
	LogPath string
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

func LoadEnv() Conf {
	C = Conf{
		Server: ServerConf{
			Host: os.Getenv("SERVER_HOST"),
			Port: os.Getenv("SERVER_PORT"),
		},
		MySQL: MySQLConf{
			DSN:         os.Getenv("MYSQL_DSN"),
			Active:      utils.StringToInt(os.Getenv("MYSQL_ACTIVE")),
			Idle:        utils.StringToInt(os.Getenv("MYSQL_IDLE")),
			IdleTimeout: time.Second * time.Duration(utils.StringToInt(os.Getenv("MYSQL_IDLE_TIMEOUT"))),
		},
		Redis: RedisConf{
			Host:            os.Getenv("REDIS_HOST"),
			Port:            os.Getenv("REDIS_PORT"),
			MaxIdle:         utils.StringToInt(os.Getenv("REDIS_MAX_IDLE")),
			MaxActive:       utils.StringToInt(os.Getenv("REDIS_MAX_ACTIVE")),
			IdleTimeout:     time.Second * time.Duration(utils.StringToInt(os.Getenv("REDIS_IDLE_TIMEOUT"))),
			MaxConnLifetime: time.Second * time.Duration(utils.StringToInt(os.Getenv("REDIS_MAX_CONN_LIFE_TIME"))),
		},
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
