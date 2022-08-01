package repository

import (
	"errors"
	"transaction/db"
	"transaction/library"
	model "transaction/module/models"
	"transaction/util"

	"github.com/gin-gonic/gin"
)

type ITranReferences interface {
	FindTransactionsInDatabase(c *gin.Context) error
}
type TranReferences struct {
	ITranReferences
	Transaction  *model.Transaction
	Transactions *[]model.Transaction
}

func (tr *TranReferences) FindTransactionsInDatabase() (err error) {
	err = db.GetGormDB().Table(library.TB_TRANSACTIONS).Find(&tr.Transactions).Error
	return
}

func BeginTransaction(transac *model.Transaction) error {
	var (
		payerAccount = &model.Account{}
		payeeAccount = &model.Account{}
		tx           = db.GetGormDB().Begin()
	)

	util.ValidateTransacion(tx, tx.Table(library.TB_ACCOUNTS).Where("id = ?", transac.IdPayer).Find(payerAccount).Error)

	util.ValidateTransacion(tx, tx.Table(library.TB_ACCOUNTS).Where("id = ?", transac.IdPayee).Find(payeeAccount).Error)

	if payerAccount.Balance < transac.Value {
		tx.Rollback()
		return errors.New("insuficient Balance")
	}

	payerAccount.Balance = payerAccount.Balance - transac.Value
	payeeAccount.Balance = payeeAccount.Balance + transac.Value

	util.ValidateTransacion(tx, tx.Table(library.TB_ACCOUNTS).Where("id = ?", transac.IdPayer).Updates(payerAccount).Error)

	util.ValidateTransacion(tx, tx.Table(library.TB_ACCOUNTS).Where("id = ?", transac.IdPayee).Updates(payeeAccount).Error)

	util.ValidateTransacion(tx, tx.Table(library.TB_TRANSACTIONS).Create(transac).Error)

	tx.Commit()

	return nil
}
