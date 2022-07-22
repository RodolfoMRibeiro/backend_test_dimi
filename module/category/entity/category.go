package entity

import "transaction/module/user/entity"

type Category struct {
	Id    int `gorm:"primaryKey"`
	Name  string
	Users []entity.User `gorm:"foreignKey:id_category"`
}
