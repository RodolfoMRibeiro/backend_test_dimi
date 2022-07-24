package entity

import "transaction/module/transaction/entity"

type Account struct {
	Id           int    `gorm:"primaryKey"`
	CpfCnpj      string `gorm:"type:varchar(14)"`
	Balance      int
	TransacPayer entity.Transaction `gorm:"foreignKey:IdPayer"`
	TransacPayee entity.Transaction `gorm:"foreignKey:IdPayee"`
}
