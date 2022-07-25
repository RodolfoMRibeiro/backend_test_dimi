package entity

import (
	"transaction/module/transaction/entity"

	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	CpfCnpj      string             `json:"cpf_cnpj" gorm:"type:varchar(14)"`
	Balance      int                `json:"balance"`
	TransacPayer entity.Transaction `json:"transac_payer" gorm:"foreignKey:IdPayer"`
	TransacPayee entity.Transaction `json:"transac_payee" gorm:"foreignKey:IdPayee"`
}

type Tabler interface {
	TableName() string
}

func (Account) TableName() string {
	return "tb_accounts"
}
