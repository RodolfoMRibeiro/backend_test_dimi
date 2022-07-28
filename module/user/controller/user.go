package controller

import (
	"net/http"
	"transaction/db"

	"transaction/module/account/controller"
	entity_user "transaction/module/user/entity"
	service_user "transaction/module/user/service"
	"transaction/util"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var newUser *entity_user.User
	if containsError(c, c.BindJSON(&newUser)) && checkEmailAndCpf_Cnpf(newUser) {
		service_user.AddUserToDatabase(c, newUser)
	}
}

func FindUser(c *gin.Context) {
	var newUsers []entity_user.User
	if err := db.GetGormDB().Find(&newUsers).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	for i := 0; i < len(newUsers); i++ {
		newUsers[i].Account = controller.GetAccountsFromUser(newUsers[i].CpfCnpj)
	}
	c.JSON(http.StatusFound, newUsers)
}

func UploadUser(c *gin.Context) {
	var newUser *entity_user.User
	if containsError(c, c.BindJSON(&newUser)) && checkEmailAndCpf_Cnpf(newUser) {
		service_user.UpdateUserInDatabase(c, newUser)
	}
}

func DeleteUser(c *gin.Context) {
	c.IndentedJSON(http.StatusNotFound, "Sorry, but this method hasn't been developed yet")
}

// -------------------------------------------< Aux funcs >------------------------------------------- \\

func containsError(c *gin.Context, err error) bool {
	if err != nil {
		c.IndentedJSON(http.StatusNotAcceptable, "wrong data inserted") // 406
		return false
	}
	return true
}

func checkEmailAndCpf_Cnpf(u *entity_user.User) (boolean bool) {
	cpf_cnpj, ok := util.VerifyingCPForCNPJ(util.LetOnlyNumbers(util.TrimAllSpacesInString(u.CpfCnpj)))
	if ok && util.IsEmailValid(util.TrimAllSpacesInString(u.Email)) {
		boolean = true
		u.CpfCnpj = cpf_cnpj
	}
	return
}
