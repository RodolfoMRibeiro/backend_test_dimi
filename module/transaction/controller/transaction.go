package controller

import (
	"net/http"
	"transaction/db"
	"transaction/module/account/entity"
	entity_transaction "transaction/module/transaction/entity"

	"github.com/gin-gonic/gin"
)

//sei que está errado não realizar a consulta no banco, refarei depois

func CreateTransaction(c *gin.Context) {
	var NewTransaction *entity_transaction.Transaction
	if err := c.BindJSON(&NewTransaction); err != nil {
		c.IndentedJSON(http.StatusNotAcceptable, "wrong data inserted") // 406
		return
	}

	NewTransaction.ValidateTransaction()
	//arrumar essa porcaria de 1 depois --> clean code prega a destruição de magic numbers
	if NewTransaction.IdPayer != NewTransaction.IdPayee && NewTransaction.IdStatus != 1 {

		beginTransaction(NewTransaction)

		c.JSON(http.StatusOK, gin.H{"New user registred": NewTransaction})
	}
}

func FindTransaction(c *gin.Context) {
	var NewTransactions *[]entity_transaction.Transaction

	if err := db.DB.Find(&NewTransactions).Error; err != nil {
		c.IndentedJSON(http.StatusNoContent, "could not find the transaction")
		return
	}

	c.JSON(http.StatusOK, NewTransactions)
}

//Perdão por esse código imenso e com esse bando de validações que ferem os princípios do código limpo
func beginTransaction(transac *entity_transaction.Transaction) error {
	var (
		payerAccount = &entity.Account{}
		payeeAccount = &entity.Account{}
	)
	tx := db.DB.Begin()

	if err := tx.Table("tb_accounts").Where("id = ?", transac.IdPayer).Find(payerAccount).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Table("tb_accounts").Where("id = ?", transac.IdPayee).Find(payeeAccount).Error; err != nil {
		tx.Rollback()
		return err
	}

	payerAccount.Balance = payerAccount.Balance - transac.Value
	payeeAccount.Balance = payeeAccount.Balance + transac.Value

	if err := tx.Table("tb_accounts").Where("id = ?", transac.IdPayer).Updates(payerAccount).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Table("tb_accounts").Where("id = ?", transac.IdPayee).Updates(payeeAccount).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Table("tb_transactions").Create(transac).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}
