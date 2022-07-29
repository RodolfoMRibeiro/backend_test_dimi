package controller

import (
	"net/http"
	"transaction/db"

	model "transaction/module/models"
	service "transaction/module/services"
	"transaction/util"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var newUser *model.User
	if util.ContainsError(c.BindJSON(&newUser)) && checkEmailAndCpf_Cnpf(newUser) {
		service.AddUserToDatabase(c, newUser)
	}
}

func FindUser(c *gin.Context) {
	var newUsers []model.User
	if err := db.GetGormDB().Find(&newUsers).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	for i := 0; i < len(newUsers); i++ {
		newUsers[i].Account = GetAccountsFromUser(newUsers[i].CpfCnpj)
	}
	c.JSON(http.StatusFound, newUsers)
}

func UploadUser(c *gin.Context) {
	var newUser *model.User
	if util.ContainsError(c.BindJSON(&newUser)) && checkEmailAndCpf_Cnpf(newUser) {
		service.UpdateUserInDatabase(c, newUser)
	}
}

func DeleteUser(c *gin.Context) {
	c.IndentedJSON(http.StatusNotFound, "Sorry, but this method hasn't been developed yet")
}

// -------------------------------------------< Aux funcs >------------------------------------------- \\

func checkEmailAndCpf_Cnpf(u *model.User) (boolean bool) {
	cpf_cnpj, ok := util.VerifyingCPForCNPJ(util.LetOnlyNumbers(util.TrimAllSpacesInString(u.CpfCnpj)))
	if ok && util.IsEmailValid(util.TrimAllSpacesInString(u.Email)) {
		boolean = true
		u.CpfCnpj = cpf_cnpj
	}
	return
}
