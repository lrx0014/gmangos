package dao

import (
	"database/sql"
	"github.com/gomodule/redigo/redis"
	"gmangos/src/libs/config"
	mysqlBuilder "gmangos/src/libs/mysql"
	redisBuilder "gmangos/src/libs/redis"
)

type Dao struct {
	db    *sql.DB
	redis *redis.Pool
}

func New(c *config.Conf) *Dao {
	return &Dao{
		db:    mysqlBuilder.New(c.MySQL),
		redis: redisBuilder.New(c.Redis),
	}
}
