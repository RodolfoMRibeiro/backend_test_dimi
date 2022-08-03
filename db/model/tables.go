package model

import "transaction/library"

type Account struct {
	Id           int         `json:"id" gorm:"primaryKey"`
	CpfCnpj      string      `json:"cpf_cnpj" gorm:"type:varchar(14)"`
	Balance      int         `json:"balance"`
	TransacPayer Transaction `json:"transaction" gorm:"foreignKey:IdPayer;autoIncrement:false"`
}

type Category struct {
	Id    int    `gorm:"primaryKey"`
	Name  string `gorm:"type:varchar(15)"`
	Users []User `gorm:"foreignKey:IdCategory"`
}

type Status struct {
	Id          int           `gorm:"primaryKey"`
	Name        string        `gorm:"type:varchar(15)"`
	Transaction []Transaction `gorm:"foreignKey:IdStatus"`
}

type Transaction struct {
	Id       int `json:"id" gorm:"primaryKey"`
	IdPayer  int `json:"id_payer"`
	IdPayee  int `json:"id_payee"`
	IdStatus int `json:"id_status"`
	Value    int `json:"value"`
}

type User struct {
	CpfCnpj    string    `json:"cpf_cnpj" gorm:"type:varchar(14);primaryKey;autoIncrement:false"`
	Email      string    `json:"email" gorm:"type:varchar(25); unique"`
	Account    []Account `json:"account" gorm:"foreignKey:CpfCnpj;references:CpfCnpj"`
	IdCategory int       `json:"id_category"`
	FullName   string    `json:"full_name" gorm:"type:varchar(30)"`
	Password   string    `json:"password" gorm:"type:varchar(30)"`
}

type Tabler interface {
	TableName() string
}

// ---------------------------------< Methods to name database tables >--------------------------------- \\

func (Account) TableName() string {
	return library.TB_ACCOUNTS
}

func (Category) TableName() string {
	return library.TB_CATEGORIES
}

func (Status) TableName() string {
	return library.TB_STATUS
}

func (Transaction) TableName() string {
	return library.TB_TRANSACTIONS
}

func (User) TableName() string {
	return library.TB_USERS
}
