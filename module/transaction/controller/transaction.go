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

	// db.DB.Joins("FROM tb_transactions A RIGHT JOIN tb_accounts ON A.id_payer = ?").Where(NewTransaction.IdPayer).Find(&controller_account.NewAccount)

	// c.JSON(http.StatusOK, gin.H{"New user registred": controller_account.NewAccount})

}

func FindTransaction(c *gin.Context) {

	db.DB.Find(&NewTransactions)
	c.JSON(http.StatusOK, NewTransactions)
}
