package entity

import "transaction/module/transaction/entity"

type Account struct {
	Id           int                `json:"id" gorm:"primaryKey"`
	CpfCnpj      string             `json:"cpf_cnpj" gorm:"type:varchar(14)"`
	Balance      int                `json:"balance"`
	TransacPayer entity.Transaction `json:"transac_payer" gorm:"foreignKey:IdPayer"`
	TransacPayee entity.Transaction `json:"transac_payee" gorm:"foreignKey:IdPayee"`
}
