package entity

import "transaction/module/account/entity"

type User struct {
	Cpf_Cnpj    string           `gorm:"primaryKey;autoIncrement:false"`
	Account     []entity.Account `gorm:"foreignKey:Cpf_Cnpj"`
	Id_category int
	Fullname    string
	Password    string
}
