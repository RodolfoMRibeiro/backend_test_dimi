package repository

import (
	"transaction/db"
	"transaction/library"
	model "transaction/module/models"
)

type IUserReferences interface {
	AddUserToDatabase() error
	FindUsersInDatabase() error
	UpdateUserInDatabase() error
	GetUserByAccountId() error
	GetAccountsFromUser() error
}

type UserReferences struct {
	IUserReferences
	User  *model.User
	Users *[]model.User
}

func (u UserReferences) AddUserToDatabase() (err error) {
	err = db.GetGormDB().Table(library.TB_USERS).Create(&u.User).Error
	return
}

func (u *UserReferences) FindUsersInDatabase() (err error) {
	err = db.GetGormDB().Table(library.TB_USERS).Find(&u.Users).Error
	return
}

func (u UserReferences) UpdateUserInDatabase() (err error) {
	err = db.GetGormDB().Table(library.TB_USERS).Where("cpf_cnpj = ?", u.User.CpfCnpj).Updates(&u.User).Error
	return
}

func (u *UserReferences) GetAccountsFromUser() (err error) {
	for index, user := range *u.Users {
		err = db.GetGormDB().Table(library.TB_ACCOUNTS).Where("cpf_cnpj = ?", user.CpfCnpj).Find(&user.Account).Error
		(*u.Users)[index].Account = user.Account
	}
	return err
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
