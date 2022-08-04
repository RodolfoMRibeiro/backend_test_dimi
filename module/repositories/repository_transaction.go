package repository

import (
	"errors"
	"transaction/db"
	model "transaction/db/model"
	"transaction/integrations"
	"transaction/library"
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
		payerAccount   = &model.Account{}
		payeeAccount   = &model.Account{}
		startedTransac = db.GetGormDB().Begin()
	)

	util.ValidateTransacion(startedTransac, startedTransac.Table(library.TB_ACCOUNTS).Where("id = ?", transac.IdPayer).Find(payerAccount).Error)

	util.ValidateTransacion(startedTransac, startedTransac.Table(library.TB_ACCOUNTS).Where("id = ?", transac.IdPayee).Find(payeeAccount).Error)

	if payerAccount.Balance < transac.Value {
		startedTransac.Rollback()
		return errors.New("insuficient Balance")
	}

	payerAccount.Balance = payerAccount.Balance - transac.Value
	payeeAccount.Balance = payeeAccount.Balance + transac.Value

	util.ValidateTransacion(startedTransac, startedTransac.Table(library.TB_ACCOUNTS).Where("id = ?", transac.IdPayer).Updates(payerAccount).Error)

	util.ValidateTransacion(startedTransac, startedTransac.Table(library.TB_ACCOUNTS).Where("id = ?", transac.IdPayee).Updates(payeeAccount).Error)

	util.ValidateTransacion(startedTransac, startedTransac.Table(library.TB_TRANSACTIONS).Create(transac).Error)

	startedTransac.Commit()

	return nil
}

func (tr *TranReferences) ValidateTransaction() {
	var transacStatus = integrations.TransactionStatus{}
	transacStatus.ConnectWithExternalAPI()
	if transacStatus.Authorization {

		tr.Transaction.IdStatus = library.AUTHORIZED_TRANSACTION_STATUS
	} else {
		tr.Transaction.IdStatus = library.DENIED_TRANSACTION_STATUS
	}
}
