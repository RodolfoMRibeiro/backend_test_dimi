package entity

type Transaction struct {
	Id       int `json:"id" gorm:"primaryKey"`
	IdPayer  int `json:"id_payer"`
	IdPayee  int `json:"id_payee"`
	IdStatus int `json:"id_status"`
	Value    int `json:"value"`
}

type Tabler interface {
	TableName() string
}

func (Transaction) TableName() string {
	return "tb_transactions"
}
