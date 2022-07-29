package repository

import (
	"transaction/db"
	"transaction/library"
	model "transaction/module/models"
	service "transaction/module/services"

	"github.com/gin-gonic/gin"
)

type UserReferences struct {
	User  *model.User
	Users *[]model.User
}

func (u *UserReferences) AddUserToDatabase(c *gin.Context) (err error) {
	err = db.GetGormDB().Create(&u.User).Error
	service.BadStatusReturn(c, err)
	return
}

func (u *UserReferences) FindUserInDatabase(c *gin.Context) (err error) {
	err = db.GetGormDB().Find(&u.Users).Error
	service.BadStatusReturn(c, err)
	return

}

func (u *UserReferences) UpdateUserInDatabase(c *gin.Context) (err error) {
	err = db.GetGormDB().Table(library.TB_USERS).Where("cpf_cnpj = ?", u.User.CpfCnpj).Updates(&u.User).Error
	service.BadStatusReturn(c, err)
	return
}

// ----------------------------------------< Special case >---------------------------------------- \\

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

func GetAccountsFromUser(cpf_cnpj string) []model.Account {
	var NewAccounts []model.Account = []model.Account{}
	if err := db.GetGormDB().Table(library.TB_ACCOUNTS).Where("cpf_cnpj = ?", cpf_cnpj).Find(&NewAccounts).Error; err != nil {
		return nil
	}
	return NewAccounts
}
