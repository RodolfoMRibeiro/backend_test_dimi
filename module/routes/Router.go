package routes

import (
	"github.com/gin-gonic/gin"
)

func Avaiable(r *gin.Engine) {
	user := r.Group("/user")
	UserRoutes(user)

	account := r.Group("/account")
	AccountRoutes(account)

	transaction := account.Group("/transaction")
	TransactionRoutes(transaction)
}
