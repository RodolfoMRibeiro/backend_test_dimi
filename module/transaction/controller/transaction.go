package controller

import (
	"net/http"
	"transaction/db"
	entity_transaction "transaction/module/transaction/entity"
	"transaction/util"

	"github.com/gin-gonic/gin"
)

var NewTransaction *entity_transaction.Transaction

func CreateTransaction(c *gin.Context) {

	if err := c.BindJSON(&NewTransaction); err != nil {
		c.IndentedJSON(http.StatusNotAcceptable, "wrong data inserted") // 406
		return
	}

	db.DB.Table("tb_transactions").Create(&NewTransaction)

	c.JSON(http.StatusOK, gin.H{"New user registred": NewTransaction})
}

func FindTransaction(c *gin.Context) {

	util.BadRequest(c, db.DB.Table("tb_transactions").First(&NewTransaction).Error)

	c.JSON(http.StatusOK, gin.H{"data": NewTransaction})
}
