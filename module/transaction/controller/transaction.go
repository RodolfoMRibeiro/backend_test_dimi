package controller

import (
	"net/http"
	"transaction/db"
	entity_transaction "transaction/module/transaction/entity"

	"github.com/gin-gonic/gin"
)

func CreateTransaction(c *gin.Context) {
	var NewTransaction *entity_transaction.Transaction
	if err := c.BindJSON(&NewTransaction); err != nil {
		c.IndentedJSON(http.StatusNotAcceptable, "wrong data inserted") // 406
		return
	}
	if NewTransaction.IdPayer != NewTransaction.IdPayee {
		NewTransaction.ValidateTransaction()
		db.DB.Table("tb_transactions").Create(&NewTransaction)
		c.JSON(http.StatusOK, gin.H{"New user registred": NewTransaction})
	}
}
func FindTransaction(c *gin.Context) {
	var NewTransactions *[]entity_transaction.Transaction
	db.DB.Find(&NewTransactions)
	c.JSON(http.StatusOK, NewTransactions)
}
