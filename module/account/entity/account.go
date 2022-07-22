package entity

type Account struct {
	Id       int `gorm:"primaryKey"`
	Cpf_Cnpj string
	Balance  int
}
