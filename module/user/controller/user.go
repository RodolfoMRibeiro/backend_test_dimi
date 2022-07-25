package controller

import (
	"net/http"
	"transaction/db"

	// controller_transaction "transaction/module/transaction/controller"
	entity_user "transaction/module/user/entity"
	"transaction/util"

	"github.com/gin-gonic/gin"
)

var newUser *entity_user.User

func CreateUser(c *gin.Context) {

	if err := c.BindJSON(&newUser); err != nil {
		c.IndentedJSON(http.StatusNotAcceptable, "wrong data inserted") // 406
		return
	}

	db.DB.Table("tb_users").Create(&newUser)

	c.JSON(http.StatusOK, gin.H{"New user registred": newUser})
}

func FindUser(c *gin.Context) {

	util.BadRequest(c, db.DB.Table("tb_users").First(&newUser).Error)

	c.JSON(http.StatusOK, gin.H{"data": newUser})
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
