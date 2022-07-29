package service

import (
	"net/http"
	"transaction/db"
	"transaction/library"

	model "transaction/module/models"

	"github.com/gin-gonic/gin"
)

func AddUserToDatabase(c *gin.Context, u *model.User) {
	if err := db.GetGormDB().Create(&u).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusCreated, u)
}

func FindUserInDatabase(c *gin.Context, us *[]model.User) {
	if err := db.GetGormDB().Find(&us).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusFound, us)
}

func UpdateUserInDatabase(c *gin.Context, u *model.User) {
	if err := db.GetGormDB().Table(library.TB_USERS).Where("cpf_cnpj = ?", u.CpfCnpj).Updates(&u).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, u)
}

func GetUserByAccountId(id int) (model.User, error) {
	var newUser = &model.User{}
	account, err := GetAccountById(id)

	if err != nil {
		return *newUser, err
	}

	if err := db.GetGormDB().Table(library.TB_USERS).Where("cpf_cnpj = ?", account.CpfCnpj).First(newUser).Error; err != nil {
		return *newUser, err
	}
	return *newUser, nil
}
