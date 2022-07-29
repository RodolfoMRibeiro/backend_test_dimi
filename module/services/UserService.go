package service

import (
	"net/http"
	model "transaction/module/models"
	"transaction/util"

	"github.com/gin-gonic/gin"
)

func CheckEmailAndCpf_Cnpf(u *model.User) (boolean bool) {
	cpf_cnpj, ok := util.VerifyingCPForCNPJ(util.LetOnlyNumbers(util.TrimAllSpacesInString(u.CpfCnpj)))
	if ok && util.IsEmailValid(util.TrimAllSpacesInString(u.Email)) {
		boolean = true
		u.CpfCnpj = cpf_cnpj
	}
	return
}

func CheckCPForCPNJ(a *model.Account) (boolean bool) {
	if cpfORcnpj, ok := util.VerifyingCPForCNPJ(util.LetOnlyNumbers(util.TrimAllSpacesInString(a.CpfCnpj))); ok {
		boolean = true
		a.CpfCnpj = cpfORcnpj
	}
	return
}

//

func BadStatusReturn(c *gin.Context, err error) {
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
	}
}

func FoundOrNotStatusReturn(err error, c *gin.Context, obj interface{}) {
	if err != nil {
		BadStatusReturn(c, err)
	} else {
		c.JSON(http.StatusFound, obj)
	}
}

//

func RegistredAccountStatus(err error, c *gin.Context, obj interface{}) {
	if err != nil {
		BadStatusReturn(c, err)
	} else {
		c.JSON(http.StatusOK, gin.H{"New transaction registred": obj})
	}
}
