package routes

import (
	controller "transaction/module/controllers"

	"github.com/gin-gonic/gin"
)

func AccountRoutes(account *gin.RouterGroup) {
	account.POST("/", controller.CreateAccount)
	account.GET("/", controller.FindAccount)
	account.PUT("/", controller.UpdateAccount)
	account.DELETE("/", controller.DeleteAccount)
}
