package entity

import "transaction/module/transaction/entity"

type Status struct {
	Id          int `gorm:"primaryKey"`
	Name        string
	Transaction []entity.Transaction `gorm:"foreignKey:IdStatus"`
}
