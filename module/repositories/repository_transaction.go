package repository

import (
	"errors"
	"transaction/db"
	"transaction/library"
	model "transaction/module/models"

	"github.com/gin-gonic/gin"
)

type ITranReferences interface {
	FindTransactionsInDatabase(c *gin.Context) error
}
type TranReferences struct {
	Transaction  *model.Transaction
	Transactions *[]model.Transaction
}

func (tr *TranReferences) FindTransactionsInDatabase(c *gin.Context) (err error) {
	err = db.GetGormDB().Find(&tr.Transactions).Error
	return
}

func BeginTransaction(transac *model.Transaction) error {
	var (
		payerAccount = &model.Account{}
		payeeAccount = &model.Account{}
		tx           = db.GetGormDB().Begin()
	)

	if err := tx.Table(library.TB_ACCOUNTS).Where("id = ?", transac.IdPayer).Find(payerAccount).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Table(library.TB_ACCOUNTS).Where("id = ?", transac.IdPayee).Find(payeeAccount).Error; err != nil {
		tx.Rollback()
		return err
	}

	if payerAccount.Balance < transac.Value {
		tx.Rollback()
		return errors.New("insuficient Balance")
	}
	payerAccount.Balance = payerAccount.Balance - transac.Value
	payeeAccount.Balance = payeeAccount.Balance + transac.Value

	if err := tx.Table(library.TB_ACCOUNTS).Where("id = ?", transac.IdPayer).Updates(payerAccount).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Table(library.TB_ACCOUNTS).Where("id = ?", transac.IdPayee).Updates(payeeAccount).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Table(library.TB_TRANSACTIONS).Create(transac).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}
