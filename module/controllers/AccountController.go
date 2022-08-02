package controller

import (
	repository "transaction/module/repositories"
	service "transaction/module/services"
	"transaction/util"

	"github.com/gin-gonic/gin"
)

func CreateAccount(c *gin.Context) {
	var new repository.AccoReferences

	if !util.ContainsError(c.BindJSON(&new.Account)) {
		err := new.AddAccountToDatabase()
		service.CreateOrNotStatusReturn(err, c, new.Account)
	}
}

func FindAccount(c *gin.Context) {
	var registered repository.AccoReferences

	err := registered.FindAccountsInDatabase()
	service.FoundOrNotStatusReturn(err, c, registered.Accounts)
}

func UpdateAccount(c *gin.Context) {
	var registered repository.AccoReferences

	if !util.ContainsError(c.BindJSON(&registered.Account)) {
		err := registered.UpdateAccountInDatabase()
		service.ModifiedOrNotStatusReturn(err, c, registered.Account)
	}
}

func DeleteAccount(c *gin.Context) {
	var registered repository.AccoReferences

	if !util.ContainsError(c.BindJSON(&registered.Account)) {
		err := registered.DeleteAccountInDatabase()
		service.DeleteOrNotStatusReturn(err, c, registered.Account)
	}
}

func DeleteAccountsByCpf_Cnpj(c *gin.Context, cpf_cnpj string) {
	var registered *repository.AccoReferences

	registered.Account.CpfCnpj = cpf_cnpj
	if service.CheckCPForCPNJ(registered.Account) && !util.ContainsError(c.BindJSON(&registered.Account)) {
		repository.DeleteByCpf_Cnpj(c, registered.Account)
	}
}
