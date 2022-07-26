package controller

import (
	"net/http"
	"transaction/db"
	entity_account "transaction/module/account/entity"

	"github.com/gin-gonic/gin"
)

func CreateAccount(c *gin.Context) {
	var NewAccount *entity_account.Account = &entity_account.Account{}
	if err := c.BindJSON(NewAccount); err != nil {
		c.IndentedJSON(http.StatusNotAcceptable, "wrong data inserted") // 406
		return
	}

	if err := db.DB.Table("tb_accounts").Create(NewAccount).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"New user registred": NewAccount})
}

func FindAccount(c *gin.Context) {
	var NewAccounts *[]entity_account.Account = &[]entity_account.Account{}

	if err := db.DB.Find(NewAccounts).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, NewAccounts)
}

func UpdateAccount(c *gin.Context) {
	var NewAccount *entity_account.Account = &entity_account.Account{}

	if err := c.BindJSON(NewAccount); err != nil {
		c.IndentedJSON(http.StatusNotAcceptable, "wrong data inserted") // 406
		return
	}

	if err := db.DB.Table("tb_accounts").Where("id = ?", NewAccount.Id).Updates(NewAccount).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
}

func DeleteAccount(c *gin.Context) {
	var NewAccount *entity_account.Account = &entity_account.Account{}

	if err := c.BindJSON(NewAccount); err != nil {
		c.IndentedJSON(http.StatusNotAcceptable, "wrong data inserted") // 406
		return
	}

	if err := db.DB.Table("tb_accounts").Where("id = ?", NewAccount.Id).Delete(NewAccount).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"Account deleted": NewAccount})
}

func DeleteAccountsByCpf_Cnpj(c *gin.Context, cpf_cnpj string) {
	var NewAccount *entity_account.Account = &entity_account.Account{}
	NewAccount.CpfCnpj = cpf_cnpj

	if err := db.DB.Table("tb_accounts").Where("cpf_cnpj = ?", cpf_cnpj).Delete(NewAccount).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"Account deleted": NewAccount})
}
