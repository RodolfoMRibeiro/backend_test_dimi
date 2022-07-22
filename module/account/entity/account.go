package entity

import "transaction/module/transaction/entity"

type Account struct {
	Id           int `gorm:"primaryKey"`
	Cpf_Cnpj     string
	Balance      int
	Transactions []entity.Transaction `gorm:"foreignKey:Id_payer;foreignKey:Id_payee"`
}
