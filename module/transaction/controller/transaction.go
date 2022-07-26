package controller

import (
	"net/http"
	"transaction/db"
	controller_account "transaction/module/account/controller"
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

	// db.DB.Where().Create(&NewTransaction)

	db.DB.Joins("FROM tb_transactions A RIGHT JOIN tb_accounts ON A.id_payer = ?").Where(NewTransaction.IdPayer).Find(&controller_account.NewAccount)
	// .Joins("RIGHT JOIN tb_accounts.cpf_cnpj ON tb_users.cpf_cnpj = ?", NewTransaction.IdPayer).Find(&NewTransaction)
	// db.DB.Table("tb_accounts").Where("id = ?", NewTransaction.).find(&status).Err
	// fmt.Println(controller_account.NewAccount.CpfCnpj)

	// db.DB.Table("tb_transactions")

	c.JSON(http.StatusOK, gin.H{"New user registred": controller_account.NewAccount})
}

func FindTransaction(c *gin.Context) {

	db.DB.Find(&NewTransactions)
	c.JSON(http.StatusOK, NewTransactions)
}
