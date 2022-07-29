package model

import (
	"transaction/integrations"
	"transaction/library"
)

type Transaction struct {
	Id       int `json:"id" gorm:"primaryKey"`
	IdPayer  int `json:"id_payer"`
	IdPayee  int `json:"id_payee"`
	IdStatus int `json:"id_status"`
	Value    int `json:"value"`
}

func (Transaction) TableName() string {
	return library.TB_TRANSACTIONS
}

func (t *Transaction) ValidateTransaction() {
	var transacStatus = integrations.TransactionStatus{}
	transacStatus.ConnectWithExternalAPI()
	if transacStatus.Authorization {
		t.IdStatus = 1
	} else {
		t.IdStatus = 2
	}
}
