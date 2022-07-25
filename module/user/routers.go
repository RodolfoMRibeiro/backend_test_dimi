package routes

import (
	module_user "transaction/module/user/controller"

	"github.com/gin-gonic/gin"
)

func UserRoutes(user *gin.RouterGroup) {
	user.POST("/", module_user.CreateUser)
	user.GET("/", module_user.FindUser)
	user.PUT("/", module_user.UploadUser)
	user.DELETE("/", module_user.DeleteUser)
}
