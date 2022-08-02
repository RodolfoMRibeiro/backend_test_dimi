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
	var registered repository.UserReferences

	err1 := registered.FindUsersInDatabase()
	err2 := registered.GetAccountsFromUser()

	fmt.Println(err1, err2)
	service.FoundOrNotStatusReturn(err2, c, registered.Users)

}

func UploadUser(c *gin.Context) {
	var registered repository.UserReferences

	if !util.ContainsError(c.BindJSON(&registered.User)) && service.CheckEmailAndCpf_Cnpf(registered.User) {
		err := registered.UpdateUserInDatabase()
		service.ModifiedOrNotStatusReturn(err, c, registered.User)
	}
}

func DeleteUser(c *gin.Context) {
	var registered repository.UserReferences
	if !util.ContainsError(c.BindJSON(&registered.User)) {
		err := registered.DeleteUserInDatabase()
		service.DeleteOrNotStatusReturn(err, c, registered.User)
	}

}
