package entity

import (
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
	return "tb_categories"
}
