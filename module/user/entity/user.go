package entity

import (
	"transaction/module/account/entity"
)

type User struct {
	CpfCnpj    string           `json:"cpf_cnpj" gorm:"type:varchar(14);primaryKey;autoIncrement:false"`
	Account    []entity.Account `json:"account" gorm:"foreignKey:CpfCnpj"`
	IdCategory int              `json:"id_category"`
	FullName   string           `json:"full_name"`
	Password   string           `json:"password"`
}
