package controller

import (
	"fmt"

	repository "transaction/module/repositories"
	service "transaction/module/services"
	"transaction/util"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var new repository.UserReferences

	if !util.ContainsError(c.BindJSON(&new.User)) && service.CheckEmailAndCpf_Cnpf(new.User) {
		err := new.AddUserToDatabase()
		service.FoundOrNotStatusReturn(err, c, new.User)
	}
}

func FindUser(c *gin.Context) {
	var registred repository.UserReferences

	err1 := registred.FindUsersInDatabase()
	err2 := registred.GetAccountsFromUser()

	fmt.Println(err1, err2)
	service.FoundOrNotStatusReturn(err2, c, registred.Users)

}

func UploadUser(c *gin.Context) {
	var registred repository.UserReferences

	if !util.ContainsError(c.BindJSON(&registred.User)) && service.CheckEmailAndCpf_Cnpf(registred.User) {
		err := registred.UpdateUserInDatabase()
		service.ModifiedOrNotStatusReturn(err, c, registred.User)
	}
}

func DeleteUser(c *gin.Context) {
	var registred repository.UserReferences
	if !util.ContainsError(c.BindJSON(&registred.User)) {
		err := registred.DeleteUserInDatabase()
		service.DeleteOrNotStatusReturn(err, c, registred.User)
	}

}
