package entity

type Transaction struct {
	Id       int `gorm:"primaryKey"`
	IdPayer  int
	IdPayee  int
	IdStatus int
	Value    int
}
