package controller

import (
	"net/http"
	"transaction/db"
	"transaction/library"
	model "transaction/module/models"
	service "transaction/module/services"
	"transaction/util"

	"github.com/gin-gonic/gin"
)

func CreateAccount(c *gin.Context) {
	var NewAccount *model.Account
	if !util.ContainsError(c.BindJSON(&NewAccount)) {
		service.AddAccountToDatabase(c, NewAccount)
	}
	c.JSON(http.StatusOK, gin.H{"New transaction registred": NewAccount})
}

func FindAccount(c *gin.Context) {
	var NewAccounts *[]model.Account
	service.FindAccountInDatabase(c, NewAccounts)
}

func UpdateAccount(c *gin.Context) {
	var NewAccount *model.Account
	if !util.ContainsError(c.BindJSON(&NewAccount)) {
		service.UpdateAccountInDatabase(c, NewAccount)
	}
}

func DeleteAccount(c *gin.Context) {
	var NewAccount *model.Account
	if !util.ContainsError(c.BindJSON(&NewAccount)) {
		service.DeleteAccountInDatabase(c, NewAccount)
	}
}

func DeleteAccountsByCpf_Cnpj(c *gin.Context, cpf_cnpj string) {
	var NewAccount *model.Account
	NewAccount.CpfCnpj = cpf_cnpj
	if checkCPForCPNJ(NewAccount) && !util.ContainsError(c.BindJSON(&NewAccount)) {
		service.DeleteByCpf_Cnpj(c, NewAccount)
	}
}

// -------------------------------------------< Aux funcs >------------------------------------------- \\

func checkCPForCPNJ(a *model.Account) (boolean bool) {
	if cpfORcnpj, ok := util.VerifyingCPForCNPJ(util.LetOnlyNumbers(util.TrimAllSpacesInString(a.CpfCnpj))); ok {
		boolean = true
		a.CpfCnpj = cpfORcnpj
	}
	return
}

func GetAccountsFromUser(cpf_cnpj string) []model.Account {
	var NewAccounts []model.Account = []model.Account{}
	if err := db.GetGormDB().Table(library.TB_ACCOUNTS).Where("cpf_cnpj = ?", cpf_cnpj).Find(&NewAccounts).Error; err != nil {
		return nil
	}
	return NewAccounts
}
