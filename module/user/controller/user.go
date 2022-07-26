package controller

import (
	"fmt"
	"net/http"
	"transaction/db"

	// controller_transaction "transaction/module/transaction/controller"
	entity_user "transaction/module/user/entity"
	"transaction/util"

	"github.com/gin-gonic/gin"
)

var newUser *entity_user.User
var newUsers *[]entity_user.User

func CreateUser(c *gin.Context) {

	if err := c.BindJSON(&newUser); err != nil {
		c.IndentedJSON(http.StatusNotAcceptable, "wrong data inserted") // 406
		return
	}

	// The email could even be an primary key for being unique, but it would break either some
	// future features or destroy the normal forms (3ºth one in special [parcial dependecies])
	if _, ok := util.VerifyingCPForCNPJ(util.LetOnlyNumbers(util.TrimAllSpacesInString(newUser.CpfCnpj))); ok {
		err := db.DB.Table("tb_users").Find(&newUser).Error
		fmt.Println("não passou do email validation")

		if err == nil && util.IsEmailValid(util.TrimAllSpacesInString(newUser.Email)) {
			fmt.Println("passou do validation")

			db.DB.Table("tb_users").Create(&newUser)

			c.JSON(http.StatusOK, gin.H{"New user registred": newUser})
		}
	}

}

func FindUser(c *gin.Context) {

	db.DB.Find(&newUsers)
	c.JSON(http.StatusOK, newUsers)
}

func UploadUser(c *gin.Context) {

	if err := c.BindJSON(&newUser); err != nil {
		c.IndentedJSON(http.StatusNotAcceptable, "wrong data inserted") // 406
		return
	}

	// db.DB.Table("tb_users").Where("cpf_cnpj = ?", c.Param("cpf_cnpj")).Updates(&newUser) --> caso queira usar a validação pelo link
	db.DB.Table("tb_users").Where("cpf_cnpj = ?", newUser.CpfCnpj).Updates(&newUser)
}

func DeleteUser(c *gin.Context) {

	if err := c.BindJSON(&newUser); err != nil {
		c.IndentedJSON(http.StatusNotAcceptable, "wrong data inserted") // 406
		return
	}

	c.JSON(http.StatusNotAcceptable, gin.H{"Warning": "contact our reponsible sector to accomplish the action"})
}
