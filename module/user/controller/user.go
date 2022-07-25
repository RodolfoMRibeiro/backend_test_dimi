package controller

import (
	"net/http"
	"transaction/db"
	entity_user "transaction/module/user/entity"
	"transaction/util"

	"github.com/gin-gonic/gin"
)

var user *entity_user.User
var newUser *entity_user.User

func CreateUser(c *gin.Context) {

	err := c.ShouldBind(&newUser)
	util.BadRequest(c, err)

	// entity_user.AddToDataBase(&newUser)
	db.DB.Create(&newUser)

	c.JSON(http.StatusOK, gin.H{"New user registred": newUser})
}

func FindUser(c *gin.Context) {

	err := db.DB.Where("id = ?", c.Param("id")).First(&newUser).Error
	util.BadRequest(c, err)

	c.JSON(http.StatusOK, gin.H{"data": newUser})
}

func UploadUser(c *gin.Context) {

	err := db.DB.Where("id = ?", c.Param("id")).First(&user).Error
	util.BadRequest(c, err)

	err1 := c.ShouldBind(&newUser)
	util.BadRequest(c, err1)

	if db.DB.Model(&user).Updates(&newUser).RowsAffected == 0 {
		db.DB.Create(&user)
	}
}

func DeleteUser(c *gin.Context) {

	err := db.DB.Where("id = ?", c.Param("id")).First(&user).Error
	util.BadRequest(c, err)

	db.DB.Delete(&user)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
