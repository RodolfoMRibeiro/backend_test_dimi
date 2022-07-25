package routes

import (
	module_account "transaction/module/account/controller"

	"github.com/gin-gonic/gin"
)

func AccountRoutes(account *gin.RouterGroup) {
	account.POST("/", module_account.CreateAccount)
	account.GET("/", module_account.FindAccount)
	account.PUT("/", module_account.UpdateAccount)
	account.DELETE("/", module_account.DeleteAccount)
}
