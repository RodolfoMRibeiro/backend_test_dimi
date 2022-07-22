package entity

import "transaction/module/transaction/entity"

type Account struct {
	Id           int `gorm:"primaryKey"`
	CpfCnpj      string
	Balance      int
	Transactions []entity.Transaction `gorm:"foreignKey:IdPayer;foreignKey:IdPayee"`
}
