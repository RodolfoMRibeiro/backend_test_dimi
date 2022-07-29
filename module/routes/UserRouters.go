package routes

import (
	controller "transaction/module/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(user *gin.RouterGroup) {
	user.POST("/", controller.CreateUser)
	user.GET("/", controller.FindUser)
	user.PUT("/", controller.UploadUser)
	user.DELETE("/", controller.DeleteUser)
}
