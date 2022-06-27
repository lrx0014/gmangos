package mysql

import (
	"gmangos/src/libs/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func New(c config.MySQLConf) *gorm.DB {
	db, err := gorm.Open(mysql.Open(c.DSN), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.SetMaxIdleConns(c.Idle)
	sqlDB.SetMaxOpenConns(c.Active)
	sqlDB.SetConnMaxLifetime(c.IdleTimeout)

	return db
}
