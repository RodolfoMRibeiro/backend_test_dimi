package routes

import (
	module_user "transaction/module/user/controller"

	"github.com/gin-gonic/gin"
)

func UserRoutes(user *gin.RouterGroup) {
	user.POST("/", module_user.CreateUser)
	user.GET("/:cpf_cnpj", module_user.FindUser)
	user.PUT("/:cpf_cnpj", module_user.UploadUser)
	user.DELETE("/:cpf_cnpj", module_user.DeleteUser)
}
