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

// ----------------------------------------< Common status >---------------------------------------- \\

func BadStatusReturn(c *gin.Context, err error) {
	if err != nil {
		c.IndentedJSON(http.StatusNotAcceptable, err)
		return
	}
}

func CreateOrNotStatusReturn(err error, c *gin.Context, obj interface{}) {
	BadStatusReturn(c, err)
	c.JSON(http.StatusCreated, gin.H{"New object registred": obj})

}
func FoundOrNotStatusReturn(err error, c *gin.Context, obj interface{}) {
	BadStatusReturn(c, err)
	c.JSON(http.StatusFound, gin.H{"Registred object Found": obj})
}

func ModifiedOrNotStatusReturn(err error, c *gin.Context, obj interface{}) {
	BadStatusReturn(c, err)
	c.JSON(http.StatusOK, gin.H{"Registred object modified": obj})
}

func DeleteOrNotStatusReturn(err error, c *gin.Context, obj interface{}) {
	BadStatusReturn(c, err)
	c.JSON(http.StatusOK, gin.H{"Registred object Deleted": obj})
}
