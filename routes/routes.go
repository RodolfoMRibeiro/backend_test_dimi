package routes

import (
	"transaction/db"
	module_account "transaction/module/account/controller"
	module_transaction "transaction/module/transaction/controller"
	module_user "transaction/module/user/controller"

	"github.com/gin-gonic/gin"
)

func Init() {
	db.Load()
}

func Avaiable(r *gin.Engine) {

	user := r.Group("/user")
	userRoutes(user)

	account := r.Group("/account")
	accountRoutes(account)

	transaction := account.Group("/transaction")
	transactionRoutes(transaction)
}

func userRoutes(user *gin.RouterGroup) {

	user.POST("/", module_user.CreateUser)
	user.GET("/:id", module_user.FindUser)
	user.PUT("/:id", module_user.UploadUser)
	user.DELETE("/:id", module_user.DeleteUser)
}

func accountRoutes(account *gin.RouterGroup) {

	account.POST("/", module_account.CreateAccount)
	account.GET("/:id", module_account.FindAccount)
	account.PUT("/:id", module_account.UpdateAccount)
	account.DELETE("/:id", module_account.DeleteAccount)
}

func transactionRoutes(transaction *gin.RouterGroup) {

	transaction.POST("/", module_transaction.CreateTransaction)
	transaction.GET("/:id", module_transaction.FindTransaction)
}
