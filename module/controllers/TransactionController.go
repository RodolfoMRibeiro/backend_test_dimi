package controller

import (
	"errors"
	"net/http"
	"transaction/db"
	"transaction/library"
	model "transaction/module/models"
	service "transaction/module/services"

	"github.com/gin-gonic/gin"
)

func FindTransaction(c *gin.Context) {
	var NewTransactions *[]model.Transaction
	if err := db.GetGormDB().Find(&NewTransactions).Error; err != nil {
		c.IndentedJSON(http.StatusNoContent, "could not find the transaction")
		return
	}
	c.JSON(http.StatusOK, NewTransactions)
}

func CreateTransaction(c *gin.Context) {
	var NewTransaction *model.Transaction
	if err := c.BindJSON(&NewTransaction); err != nil {
		c.IndentedJSON(http.StatusNotAcceptable, "wrong data inserted") // 406
		return
	}
	NewTransaction.ValidateTransaction()

	if NewTransaction.IdPayer != NewTransaction.IdPayee && NewTransaction.IdStatus == library.STORE_KEEPER_STATUS && !isLojista(NewTransaction.IdPayer) {
		if err := beginTransaction(NewTransaction); err != nil {
			c.IndentedJSON(http.StatusInternalServerError, err)
			return
		}
		c.JSON(http.StatusOK, gin.H{"New user registred": NewTransaction})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Ops! something went wrong"})
	}
}

func isLojista(AccountId int) bool {
	user, _ := service.GetUserByAccountId(AccountId)
	if user.IdCategory == 1 {
		return true
	} else {
		return false
	}
}

func beginTransaction(transac *model.Transaction) error {
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
