package dao

import (
	"database/sql"
	"github.com/gomodule/redigo/redis"

	_ "github.com/go-sql-driver/mysql"
)

type Dao struct {
	db    *sql.DB
	redis *redis.Pool
}

func New() *Dao {
	return &Dao{
		db:    newDB(),
		redis: newRedis(),
	}
}

func newDB() *sql.DB {
	dbPool, err := sql.Open("", "")
	if err != nil {
		panic(err)
	}

	return dbPool
}

func newRedis() *redis.Pool {
	redisPool := &redis.Pool{}
	return redisPool
}
