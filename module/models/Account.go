package model

import (
	"transaction/library"
)

type Account struct {
	Id           int         `json:"id" gorm:"primaryKey"`
	CpfCnpj      string      `json:"cpf_cnpj" gorm:"type:varchar(14)"`
	Balance      int         `json:"balance"`
	TransacPayer Transaction `json:"transaction" gorm:"foreignKey:IdPayer;autoIncrement:false"`
}

func (Account) TableName() string {
	return library.TB_ACCOUNTS
}
