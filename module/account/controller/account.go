package controller

import (
	"net/http"
	"transaction/db"
	entity_account "transaction/module/account/entity"

	"github.com/gin-gonic/gin"
)

var NewAccount *entity_account.Account
var NewAccounts *[]entity_account.Account

func CreateAccount(c *gin.Context) {

	if err := c.BindJSON(&NewAccount); err != nil {
		c.IndentedJSON(http.StatusNotAcceptable, "wrong data inserted") // 406
		return
	}

	db.DB.Table("tb_accounts").Create(&NewAccount)

	c.JSON(http.StatusOK, gin.H{"New user registred": NewAccount})
}

func FindAccount(c *gin.Context) {

	db.DB.Find(&NewAccounts)
	c.JSON(http.StatusOK, NewAccounts)
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
