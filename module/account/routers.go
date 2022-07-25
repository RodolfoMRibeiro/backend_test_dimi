package routes

import (
	module_account "transaction/module/account/controller"

	"github.com/gin-gonic/gin"
)

func AccountRoutes(account *gin.RouterGroup) {
	account.POST("/", module_account.CreateAccount)
	account.GET("/:id", module_account.FindAccount)
	account.PUT("/:id", module_account.UpdateAccount)
	account.DELETE("/:id", module_account.DeleteAccount)
}
