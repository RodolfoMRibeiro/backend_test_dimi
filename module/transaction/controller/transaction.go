package controller

import (
	"net/http"
	"transaction/db"
	entity_transaction "transaction/module/transaction/entity"

	"github.com/gin-gonic/gin"
)

var NewTransaction *entity_transaction.Transaction

var NewTransactions *[]entity_transaction.Transaction

func CreateTransaction(c *gin.Context) {

	if err := c.BindJSON(&NewTransaction); err != nil {
		c.IndentedJSON(http.StatusNotAcceptable, "wrong data inserted") // 406
		return
	}

	if ok := db.DB.Where("tb_users.id_category = ? and tb_accounts.id = tb.transactions.id_payee", 2).Error; ok != nil {

		db.DB.Table("tb_transactions").Create(&NewTransaction)
	}

	c.JSON(http.StatusOK, gin.H{"New user registred": NewTransaction})
}

func FindTransaction(c *gin.Context) {

	db.DB.Find(&NewTransactions)
	c.JSON(http.StatusOK, NewTransactions)
}
