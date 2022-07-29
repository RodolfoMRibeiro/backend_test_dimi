package model

import (
	"transaction/library"
)

type Category struct {
	Id    int    `gorm:"primaryKey"`
	Name  string `gorm:"type:varchar(15)"`
	Users []User `gorm:"foreignKey:IdCategory"`
}

func (Category) TableName() string {
	return library.TB_CATEGORIES
}
