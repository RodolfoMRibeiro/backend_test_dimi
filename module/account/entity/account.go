package entity

import (
	"transaction/module/transaction/entity"
)

type Account struct {
	Id           int                `json:"id" gorm:"primaryKey";autoIncrement:1`
	CpfCnpj      string             `json:"cpf_cnpj" gorm:"type:varchar(14)"`
	Balance      int                `json:"balance"`
	TransacPayer entity.Transaction `json:"transaction" gorm:"foreignKey:IdPayer;autoIncrement:false"`
}

type Tabler interface {
	TableName() string
}

func (Account) TableName() string {
	return "tb_accounts"
}
