package routes

import (
	controller "transaction/module/controllers"

	"github.com/gin-gonic/gin"
)

func TransactionRoutes(transaction *gin.RouterGroup) {
	transaction.POST("/", controller.CreateTransaction)
	transaction.GET("/", controller.FindTransaction)
}
