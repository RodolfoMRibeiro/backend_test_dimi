package entity

import (
	"transaction/library"
	"transaction/module/user/entity"
)

type Category struct {
	Id    int           `gorm:"primaryKey"`
	Name  string        `gorm:"type:varchar(15)"`
	Users []entity.User `gorm:"foreignKey:IdCategory"`
}

type Tabler interface {
	TableName() string
}

func (Category) TableName() string {
	return library.TB_CATEGORIES
}
