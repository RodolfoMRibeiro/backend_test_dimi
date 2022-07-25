package routes

import (
	module_transaction "transaction/module/transaction/controller"

	"github.com/gin-gonic/gin"
)

func TransactionRoutes(transaction *gin.RouterGroup) {
	transaction.POST("/", module_transaction.CreateTransaction)
	transaction.GET("/", module_transaction.FindTransaction)
}
