package controller

import (
	"net/http"
	"transaction/db"
	entity_user "transaction/module/user/entity"
	"transaction/util"

	"github.com/gin-gonic/gin"
)

var newUser *entity_user.User

func CreateUser(c *gin.Context) {

	util.BadRequest(c, c.ShouldBind(&newUser))

	db.DB.Create(&newUser)

	c.JSON(http.StatusOK, gin.H{"New user registred": newUser})
}

func FindUser(c *gin.Context) {

	util.BadRequest(c, db.DB.Where("id = ?", c.Param("id")).First(&newUser).Error)

	c.JSON(http.StatusOK, gin.H{"data": newUser})
}

func UploadUser(c *gin.Context) {

	util.BadRequest(c, db.DB.Where("id = ?", c.Param("id")).First(&newUser).Error)

	util.BadRequest(c, c.ShouldBind(&newUser))

	if db.DB.Model(&newUser).Updates(&newUser).RowsAffected == 0 {
		db.DB.Create(&newUser)
	}
}

func DeleteUser(c *gin.Context) {

	util.BadRequest(c, db.DB.Where("id = ?", c.Param("id")).First(&newUser).Error)

	db.DB.Delete(&newUser)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
