package entity

import (
	"transaction/module/account/entity"
)

type User struct {
	CpfCnpj    string           `json:"cpf_cnpj" gorm:"type:varchar(14);primaryKey;autoIncrement:false"`
	Email      string           `json:"email" gorm:"type:varchar(25)"`
	Account    []entity.Account `json:"account" gorm:"foreignKey:CpfCnpj;autoIncrement:false"`
	IdCategory int              `json:"id_category"`
	FullName   string           `json:"full_name" gorm:"type:varchar(30)"`
	Password   string           `json:"password" gorm:"type:varchar(30)"`
}

type Tabler interface {
	TableName() string
}

func (User) TableName() string {
	return "tb_users"
}
