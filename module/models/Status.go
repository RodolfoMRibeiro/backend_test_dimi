package model

import (
	"transaction/library"
)

type Status struct {
	Id          int           `gorm:"primaryKey"`
	Name        string        `gorm:"type:varchar(15)"`
	Transaction []Transaction `gorm:"foreignKey:IdStatus"`
}

func (Status) TableName() string {
	return library.TB_STATUS
}
