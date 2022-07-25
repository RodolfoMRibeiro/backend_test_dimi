package routes

import (
	module_user "transaction/module/user/controller"

	"github.com/gin-gonic/gin"
)

func UserRoutes(user *gin.RouterGroup) {
	user.POST("/", module_user.CreateUser)
	user.GET("/:id", module_user.FindUser)
	user.PUT("/:id", module_user.UploadUser)
	user.DELETE("/:id", module_user.DeleteUser)
}
