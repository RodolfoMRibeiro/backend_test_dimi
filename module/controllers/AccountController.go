package controller

import (
	repository "transaction/module/repositories"
	service "transaction/module/services"
	"transaction/util"

	"github.com/gin-gonic/gin"
)

func CreateAccount(c *gin.Context) {
	var new *repository.AccoReferences

	if !util.ContainsError(c.BindJSON(&new.Account)) {
		err := new.AddAccountToDatabase(c)
		service.RegistredAccountStatus(err, c, new.Account)
	}

}

func FindAccount(c *gin.Context) {
	var registred *repository.AccoReferences
	err := registred.FindAccountInDatabase(c)
	service.FoundOrNotStatusReturn(err, c, registred.Accounts)
}

func UpdateAccount(c *gin.Context) {
	var registred *repository.AccoReferences

	if !util.ContainsError(c.BindJSON(&registred.Account)) {
		err := registred.UpdateAccountInDatabase(c)
		service.FoundOrNotStatusReturn(err, c, registred.Account)
	}
}

func DeleteAccount(c *gin.Context) {
	var registred *repository.AccoReferences

	if !util.ContainsError(c.BindJSON(&registred.Account)) {
		err := registred.DeleteAccountInDatabase(c)
		service.FoundOrNotStatusReturn(err, c, registred.Account)
	}
}

func DeleteAccountsByCpf_Cnpj(c *gin.Context, cpf_cnpj string) {
	var registred *repository.AccoReferences

	registred.Account.CpfCnpj = cpf_cnpj
	if service.CheckCPForCPNJ(registred.Account) && !util.ContainsError(c.BindJSON(&registred.Account)) {
		repository.DeleteByCpf_Cnpj(c, registred.Account)
	}
}
