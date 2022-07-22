package entity

type Transaction struct {
	Id        int `gorm:"primaryKey"`
	Id_payer  int
	Id_payee  int
	Id_status int
	Value     int
}
