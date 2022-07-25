package controller

import (
	"net/http"
	"transaction/db"
	entity_account "transaction/module/account/entity"
	"transaction/util"

	"github.com/gin-gonic/gin"
)

var newAccount *entity_account.Account

func CreateAccount(c *gin.Context) {

	util.BadRequest(c, c.ShouldBind(&newAccount))

	db.DB.Table("tb_accounts").Create(&newAccount)

	c.JSON(http.StatusOK, gin.H{"New user registred": newAccount})
}

func FindAccount(c *gin.Context) {

	util.BadRequest(c, db.DB.Table("tb_accounts").Where("id = ?", c.Param("id")).First(&newAccount).Error)

	c.JSON(http.StatusOK, gin.H{"data": newAccount})
}

func UpdateAccount(c *gin.Context) {

	util.BadRequest(c, db.DB.Table("tb_accounts").Where("id = ?", c.Param("id")).First(&newAccount).Error)

	util.BadRequest(c, c.ShouldBind(&newAccount))

	if db.DB.Table("tb_accounts").Model(&newAccount).Updates(&newAccount).RowsAffected == 0 {
		db.DB.Table("tb_accounts").Create(&newAccount)
	}
}

func DeleteAccount(c *gin.Context) {

	util.BadRequest(c, db.DB.Table("tb_accounts").Where("id = ?", c.Param("id")).First(&newAccount).Error)

	db.DB.Table("tb_accounts").Delete(&newAccount)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
