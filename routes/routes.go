package routes

import (
	"transaction/controller"

	"github.com/gin-gonic/gin"
)

func Init() {

}

func Avaiable(r *gin.Engine) {

	user := r.Group("/user")
	{
		user.POST("/", controller.CreateUser)
		user.GET("/:id", controller.FindUser)
		user.PUT("/:id", controller.UploadUser)
		user.DELETE("/:id", controller.DeleteUser)
	}

	account := r.Group("/account")
	{
		account.POST("/", controller.CreateAccount)
		account.GET("/:id", controller.FindAccount)
		account.PUT("/:id", controller.UpdateAccount)
		account.DELETE("/:id", controller.DeleteAccount)

		transaction := account.Group("/transaction")
		{
			transaction.POST("/", controller.CreateTransaction)
			transaction.GET("/:id", controller.FindTransaction)
		}
	}

}
