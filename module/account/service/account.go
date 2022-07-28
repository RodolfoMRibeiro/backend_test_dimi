package service

import (
	"net/http"
	"transaction/db"
	"transaction/library"
	entity_account "transaction/module/account/entity"

	"github.com/gin-gonic/gin"
)

// -----------------------------------------< feed database >----------------------------------------- \\

func AddAccountToDatabase(c *gin.Context, a *entity_account.Account) {
	if err := db.GetGormDB().Table(library.TB_ACCOUNTS).Create(&a).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusCreated, gin.H{"New Account registred": a})
}

func FindAccountInDatabase(c *gin.Context, as *[]entity_account.Account) {
	if err := db.GetGormDB().Table(library.TB_ACCOUNTS).Find(&as).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusFound, as)
}

func UpdateAccountInDatabase(c *gin.Context, a *entity_account.Account) {
	if err := db.GetGormDB().Table(library.TB_ACCOUNTS).Where("id = ?", a.Id).Updates(&a).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, a)
}

func DeleteAccountInDatabase(c *gin.Context, a *entity_account.Account) {
	if err := db.GetGormDB().Table(library.TB_ACCOUNTS).Where("cpf_cnpj = ?", a.CpfCnpj).Delete(&a).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"Account deleted": a})
}

func DeleteByCpf_Cnpj(c *gin.Context, a *entity_account.Account) {

	if err := db.GetGormDB().Table(library.TB_ACCOUNTS).Where("cpf_cnpj = ?", a.CpfCnpj).Delete(&a).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"Account deleted": a})
}

func GetAccountById(id int) (entity_account.Account, error) {
	var NewAccount *entity_account.Account
	if err := db.GetGormDB().Table(library.TB_ACCOUNTS).Where("id = ?", id).First(&NewAccount).Error; err != nil {
		return *NewAccount, err
	}
	return *NewAccount, nil
}
