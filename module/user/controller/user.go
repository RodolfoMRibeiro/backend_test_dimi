package controller

import (
	"net/http"
	"transaction/db"
	entity_user "transaction/module/user/entity"
	"transaction/util"

	"github.com/gin-gonic/gin"
)

var newUser *entity_user.User //= &entity_user.User{}

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

	util.BadRequest(c, db.DB.Table("tb_users").Where("cpf_cnpj = ?", c.Param("cpf_cnpj")).First(&newUser).Error)

	util.BadRequest(c, c.ShouldBind(&newUser))

	if db.DB.Table("tb_users").Model(&newUser).Updates(&newUser).RowsAffected == 0 {
		db.DB.Table("tb_users").Create(&newUser)
	}
}

func DeleteUser(c *gin.Context) {

	util.BadRequest(c, db.DB.Table("tb_users").Where("cpf_cnpj = ?", c.Param("cpf_cnpj")).First(&newUser).Error)

	db.DB.Table("tb_users").Delete(&newUser)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
