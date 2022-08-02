package repository

import (
	"transaction/db"
	"transaction/library"
	model "transaction/module/models"
	"transaction/util"
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

func (u UserReferences) DeleteUserInDatabase() (err error) {
	err = DeletingProcess(u.User)
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

func DeletingProcess(user *model.User) error {
	var (
		account     = &[]model.Account{}
		transaction = &[]model.Transaction{}
		tx          = db.GetGormDB().Begin()
	)
	util.ValidateTransacion(tx, tx.Table(library.TB_ACCOUNTS).Where("cpf_cnpj = ?", user.CpfCnpj).Find(account).Error)

	util.ValidateTransacion(tx, tx.Table(library.TB_TRANSACTIONS).Find(transaction).Error)
	// --------------------------------< Start deleting >-------------------------------- \\
	for index, transac := range *transaction {
		util.ValidateTransacion(tx, tx.Table(library.TB_TRANSACTIONS).Where("id_payer = ? ", (*account)[index]).Delete(transac).Error)
	}

	util.ValidateTransacion(tx, tx.Table(library.TB_ACCOUNTS).Where("cpf_cnpj = ?", user.CpfCnpj).Delete(account).Error)

	util.ValidateTransacion(tx, tx.Table(library.TB_USERS).Where("cpf_cnpj = ?", user.CpfCnpj).Delete(user).Error)

	tx.Commit()
	return nil
}
