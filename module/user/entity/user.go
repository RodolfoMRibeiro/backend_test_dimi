package entity

import (
	"encoding/json"
	"transaction/module/account/entity"
)

type User struct {
	CpfCnpj    string           `json:"cpf_cnpj" gorm:"type:varchar(14);primaryKey;autoIncrement:false"`
	Email      string           `gorm:"type:varchar(25)"`
	Account    []entity.Account `json:"account" gorm:"foreignKey:CpfCnpj"`
	IdCategory int              `json:"id_category"`
	FullName   string           `json:"full_name" gorm:"type:varchar(30)"`
	Password   string           `json:"password" gorm:"type:varchar(30)"`
}

func (u *User) SetBook(jsonElement string) {
	json.Unmarshal([]byte(jsonElement), u)
}
