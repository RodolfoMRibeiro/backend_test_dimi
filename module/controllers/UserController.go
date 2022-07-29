package controller

import (
	"net/http"

	repository "transaction/module/repositories"
	service "transaction/module/services"
	"transaction/util"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var new *repository.UserReferences

	if !util.ContainsError(c.BindJSON(&new.User)) && service.CheckEmailAndCpf_Cnpf(new.User) {
		err := new.AddUserToDatabase(c)
		service.FoundOrNotStatusReturn(err, c, new.User)
	}
}

func FindUser(c *gin.Context) {
	var registred repository.UserReferences

	err := registred.FindUsersInDatabase(c)
	service.FoundOrNotStatusReturn(err, c, registred.Users)

	for _, userRegistred := range *registred.Users {
		userRegistred.Account = repository.GetAccountsFromUser(userRegistred.CpfCnpj)
	}
	service.FoundOrNotStatusReturn(err, c, registred.Users)

}

func UploadUser(c *gin.Context) {
	var registred *repository.UserReferences

	if !util.ContainsError(c.BindJSON(&registred.User)) && service.CheckEmailAndCpf_Cnpf(registred.User) {
		err := registred.UpdateUserInDatabase(c)
		service.ModifiedOrNotStatusReturn(err, c, registred.User)
	}
}

func DeleteUser(c *gin.Context) {
	// No way to remove all database dependencies and interfer in other transactions
	c.IndentedJSON(http.StatusNotFound, "Sorry, but this method hasn't been developed yet")
}
