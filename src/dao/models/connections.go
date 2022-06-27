package models

import "time"

type Connections struct {
	Id          int64     `gorm:"column:id" db:"id" json:"id" form:"id"`                                         //主键
	Fd          int       `gorm:"column:fd" db:"fd" json:"fd" form:"fd"`                                         //fd id
	ConnectTime time.Time `gorm:"column:connect_time" db:"connect_time" json:"connect_time" form:"connect_time"` //连接创建时间
}
