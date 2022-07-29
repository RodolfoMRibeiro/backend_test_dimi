package repository

import (
	"transaction/db"
	"transaction/library"
	model "transaction/module/models"

	"github.com/gin-gonic/gin"
)

type IUserReferences interface {
	AddUserToDatabase(c *gin.Context) error
	FindUsersInDatabase(c *gin.Context) error
	UpdateUserInDatabase(c *gin.Context) error
	GetUserByAccountId(c *gin.Context) error
	GetAccountsFromUser(c *gin.Context) error
}

type UserReferences struct {
	IUserReferences
	User  *model.User
	Users *[]model.User
}

func (u *UserReferences) AddUserToDatabase(c *gin.Context) (err error) {
	err = db.GetGormDB().Create(&u.User).Error
	return
}

func (u *UserReferences) FindUsersInDatabase(c *gin.Context) (err error) {
	err = db.GetGormDB().Find(&u.Users).Error
	return
}

func (u *UserReferences) UpdateUserInDatabase(c *gin.Context) (err error) {
	err = db.GetGormDB().Table(library.TB_USERS).Where("cpf_cnpj = ?", u.User.CpfCnpj).Updates(&u.User).Error
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
