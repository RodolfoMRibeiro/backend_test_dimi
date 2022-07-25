package routes

import (
	account_routes "transaction/module/account"
	transaction_routes "transaction/module/transaction"
	user_routes "transaction/module/user"

	"github.com/gin-gonic/gin"
)

func Avaiable(r *gin.Engine) {
	user := r.Group("/user")
	user_routes.UserRoutes(user)

	account := r.Group("/account")
	account_routes.AccountRoutes(account)

	transaction := account.Group("/transaction")
	transaction_routes.TransactionRoutes(transaction)
}
