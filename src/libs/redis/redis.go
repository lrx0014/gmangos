package redis

import (
	"github.com/gomodule/redigo/redis"
	"gmangos/src/libs/config"
)

func New(c config.RedisConf) *redis.Pool {
	return &redis.Pool{
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", c.Address())
		},
		MaxIdle:         c.MaxIdle,
		MaxActive:       c.MaxActive,
		IdleTimeout:     c.IdleTimeout,
		MaxConnLifetime: c.MaxConnLifetime,
	}
}
