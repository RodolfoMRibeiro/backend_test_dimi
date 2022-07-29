package routes

import (
	"transaction/module/routes"

	"github.com/gin-gonic/gin"
)

func Avaiable(r *gin.Engine) {
	user := r.Group("/user")
	routes.UserRoutes(user)

	account := r.Group("/account")
	routes.AccountRoutes(account)

	transaction := account.Group("/transaction")
	routes.TransactionRoutes(transaction)
}
