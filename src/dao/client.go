package dao

import (
	"github.com/gomodule/redigo/redis"
	"gmangos/src/libs/config"
	mysqlBuilder "gmangos/src/libs/mysql"
	redisBuilder "gmangos/src/libs/redis"
	"gorm.io/gorm"
)

type Dao struct {
	db    *gorm.DB
	redis *redis.Pool
}

func New(c *config.Conf) *Dao {
	return &Dao{
		db:    mysqlBuilder.New(c.MySQL),
		redis: redisBuilder.New(c.Redis),
	}
}
