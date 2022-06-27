package mysql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"gmangos/src/libs/config"
)

func New(c config.MySQLConf) *sql.DB {
	db, err := sql.Open("mysql", c.DSN)
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(c.Idle)
	db.SetMaxOpenConns(c.Active)
	db.SetConnMaxLifetime(c.IdleTimeout)

	return db
}
