package model

import (
	"transaction/library"
)

type User struct {
	CpfCnpj    string    `json:"cpf_cnpj" gorm:"type:varchar(14);primaryKey;autoIncrement:false"`
	Email      string    `json:"email" gorm:"type:varchar(25); unique"`
	Account    []Account `json:"account" gorm:"foreignKey:CpfCnpj;references:CpfCnpj"`
	IdCategory int       `json:"id_category"`
	FullName   string    `json:"full_name" gorm:"type:varchar(30)"`
	Password   string    `json:"password" gorm:"type:varchar(30)"`
}

func (User) TableName() string {
	return library.TB_USERS
}
