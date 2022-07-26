package entity

import (
	"transaction/module/transaction/entity"
)

type Status struct {
	Id          int                  `gorm:"primaryKey"`
	Name        string               `gorm:"type:varchar(15)"`
	Transaction []entity.Transaction `gorm:"foreignKey:IdStatus"`
}

type Tabler interface {
	TableName() string
	FeedAssociatedTable()
}

func (Status) TableName() string {
	return "tb_status"
}
