package routes

import (
	"github.com/gin-gonic/gin"
)

func Init() {

}

func Avaiable(r *gin.Engine) {
	user := r.Group("/user")
	{
		user.POST("/")
		user.GET("/:id")
		user.PUT("/:id")
		user.DELETE("/:id")
	}

	account := r.Group("/account")
	{
		account.POST("/")
		account.GET("/:id")
		account.PUT("/:id")
		account.DELETE("/:id")

		transaction := account.Group("/transaction")
		{
			transaction.POST("/")
			transaction.GET("/:id")
		}
	}

}
