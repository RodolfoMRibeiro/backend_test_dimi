package service

import (
	"net/http"
	"transaction/db"
	"transaction/library"

	service_account "transaction/module/account/service"
	entity_user "transaction/module/user/entity"

	"github.com/gin-gonic/gin"
)

func AddUserToDatabase(c *gin.Context, u *entity_user.User) {
	if err := db.GetGormDB().Create(&u).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusCreated, u)
}

func FindUserInDatabase(c *gin.Context, us *[]entity_user.User) {
	if err := db.GetGormDB().Find(&us).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusFound, us)
}

func UpdateUserInDatabase(c *gin.Context, u *entity_user.User) {
	if err := db.GetGormDB().Table(library.TB_USERS).Where("cpf_cnpj = ?", u.CpfCnpj).Updates(&u).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, u)
}

func GetUserByAccountId(id int) (entity_user.User, error) {
	var newUser = &entity_user.User{}
	account, err := service_account.GetAccountById(id)

	if err != nil {
		return *newUser, err
	}

	if err := db.GetGormDB().Table(library.TB_USERS).Where("cpf_cnpj = ?", account.CpfCnpj).First(newUser).Error; err != nil {
		return *newUser, err
	}
	return *newUser, nil
}
