package routes

import (
	"transaction/controller"

	"github.com/gin-gonic/gin"
)

func Init() {

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

	user.POST("/", controller.CreateUser)
	user.GET("/:id", controller.FindUser)
	user.PUT("/:id", controller.UploadUser)
	user.DELETE("/:id", controller.DeleteUser)
}

func accountRoutes(account *gin.RouterGroup) {

	account.POST("/", controller.CreateAccount)
	account.GET("/:id", controller.FindAccount)
	account.PUT("/:id", controller.UpdateAccount)
	account.DELETE("/:id", controller.DeleteAccount)
}

func transactionRoutes(transaction *gin.RouterGroup) {

	transaction.POST("/", controller.CreateTransaction)
	transaction.GET("/:id", controller.FindTransaction)
}
