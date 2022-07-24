package entity

import "transaction/module/account/entity"

type User struct {
	CpfCnpj    string           `gorm:"type:varchar(14);primaryKey;autoIncrement:false"`
	Account    []entity.Account `gorm:"foreignKey:CpfCnpj"`
	IdCategory int
	Fullname   string
	Password   string
}
