package controller

import (
	"net/http"
	"transaction/db"
	entity_account "transaction/module/account/entity"
	"transaction/util"

	"github.com/gin-gonic/gin"
)

var NewAccount *entity_account.Account

func CreateAccount(c *gin.Context) {

	if err := c.BindJSON(&NewAccount); err != nil {
		c.IndentedJSON(http.StatusNotAcceptable, "wrong data inserted") // 406
		return
	}

	db.DB.Table("tb_accounts").Create(&NewAccount)

	c.JSON(http.StatusOK, gin.H{"New user registred": NewAccount})
}

func FindAccount(c *gin.Context) {

	util.BadRequest(c, db.DB.Table("tb_accounts").First(&NewAccount).Error)

	c.JSON(http.StatusOK, gin.H{"data": NewAccount})
}

func UpdateAccount(c *gin.Context) {

	if err := c.BindJSON(&NewAccount); err != nil {
		c.IndentedJSON(http.StatusNotAcceptable, "wrong data inserted") // 406
		return
	}

	db.DB.Table("tb_accounts").Where("id = ?", NewAccount.ID).Updates(&NewAccount)
}

func DeleteAccount(c *gin.Context) {

	if err := c.BindJSON(&NewAccount); err != nil {
		c.IndentedJSON(http.StatusNotAcceptable, "wrong data inserted") // 406
		return
	}

	db.DB.Table("tb_accounts").Where("id = ?", NewAccount.ID).Delete(&NewAccount)

	c.JSON(http.StatusOK, gin.H{"Account deleted": NewAccount})
}
