package controller

import (
	"net/http"
	"transaction/db"
	entity_transaction "transaction/module/transaction/entity"
	"transaction/util"

	"github.com/gin-gonic/gin"
)

var newTransaction *entity_transaction.Transaction

func CreateTransaction(c *gin.Context) {

	util.BadRequest(c, c.ShouldBind(&newTransaction))

	db.DB.Create(&newTransaction)

	c.JSON(http.StatusOK, gin.H{"New user registred": newTransaction})
}

func FindTransaction(c *gin.Context) {

	util.BadRequest(c, db.DB.Where("id = ?", c.Param("id")).First(&newTransaction).Error)

	c.JSON(http.StatusOK, gin.H{"data": newTransaction})
}
