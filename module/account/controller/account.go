package controller

import (
	"net/http"
	"transaction/db"
	"transaction/library"
	entity_account "transaction/module/account/entity"
	service_account "transaction/module/account/service"
	"transaction/util"

	"github.com/gin-gonic/gin"
)

func CreateAccount(c *gin.Context) {
	var NewAccount *entity_account.Account
	if !util.ContainsError(c.BindJSON(&NewAccount)) {
		service_account.AddAccountToDatabase(c, NewAccount)
	}
	c.JSON(http.StatusOK, gin.H{"New transaction registred": NewAccount})
}

func FindAccount(c *gin.Context) {
	var NewAccounts *[]entity_account.Account
	service_account.FindAccountInDatabase(c, NewAccounts)
}

func UpdateAccount(c *gin.Context) {
	var NewAccount *entity_account.Account
	if !util.ContainsError(c.BindJSON(&NewAccount)) {
		service_account.UpdateAccountInDatabase(c, NewAccount)
	}
}

func DeleteAccount(c *gin.Context) {
	var NewAccount *entity_account.Account
	if !util.ContainsError(c.BindJSON(&NewAccount)) {
		service_account.DeleteAccountInDatabase(c, NewAccount)
	}
}

func DeleteAccountsByCpf_Cnpj(c *gin.Context, cpf_cnpj string) {
	var NewAccount *entity_account.Account
	NewAccount.CpfCnpj = cpf_cnpj
	if checkCPForCPNJ(NewAccount) && !util.ContainsError(c.BindJSON(&NewAccount)) {
		service_account.DeleteByCpf_Cnpj(c, NewAccount)
	}
}

// -------------------------------------------< Aux funcs >------------------------------------------- \\

func checkCPForCPNJ(a *entity_account.Account) (boolean bool) {
	if cpfORcnpj, ok := util.VerifyingCPForCNPJ(util.LetOnlyNumbers(util.TrimAllSpacesInString(a.CpfCnpj))); ok {
		boolean = true
		a.CpfCnpj = cpfORcnpj
	}
	return
}

func GetAccountsFromUser(cpf_cnpj string) []entity_account.Account {
	var NewAccounts []entity_account.Account = []entity_account.Account{}
	if err := db.GetGormDB().Table(library.TB_ACCOUNTS).Where("cpf_cnpj = ?", cpf_cnpj).Find(&NewAccounts).Error; err != nil {
		return nil
	}
	return NewAccounts
}
