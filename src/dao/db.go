package dao

import (
	"gmangos/src/dao/models"
	"time"
)

func (d *Dao) SaveConnState(fd int, t time.Time) (err error) {
	connection := &models.Connections{
		Fd:          fd,
		ConnectTime: t,
	}

	return d.db.Create(connection).Error
}
