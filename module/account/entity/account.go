package entity

import (
	"transaction/library"
	"transaction/module/transaction/entity"
)

type Account struct {
	Id           int                `json:"id" gorm:"primaryKey"`
	CpfCnpj      string             `json:"cpf_cnpj" gorm:"type:varchar(14)"`
	Balance      int                `json:"balance"`
	TransacPayer entity.Transaction `json:"transaction" gorm:"foreignKey:IdPayer;autoIncrement:false"`
}

type Tabler interface {
	TableName() string
}

func (Account) TableName() string {
	return library.TB_ACCOUNTS
}
