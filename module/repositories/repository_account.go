package repository

import (
	"net/http"
	"transaction/db"
	"transaction/library"
	model "transaction/module/models"

	"github.com/gin-gonic/gin"
)

type IAccountReferences interface {
	GetAccountById(id int) (model.Account, error)
	AddAccountToDatabase() error
	FindAccountsInDatabase() error
	UpdateAccountInDatabase() error
	DeleteAccountInDatabase() error
	DeleteByCpf_Cnpj(c *gin.Context, a *model.Account) error
}

type AccoReferences struct {
	IAccountReferences
	Account  *model.Account
	Accounts *[]model.Account
}

func (ac *AccoReferences) AddAccountToDatabase() (err error) {
	err = db.GetGormDB().Table(library.TB_ACCOUNTS).Create(&ac.Account).Error
	return

	// c.JSON(http.StatusCreated, gin.H{"New Account registred": a})
}

func (ac *AccoReferences) FindAccountsInDatabase() (err error) {
	err = db.GetGormDB().Table(library.TB_ACCOUNTS).Find(&ac.Accounts).Error
	return

	// c.JSON(http.StatusFound, as)
}

func (ac *AccoReferences) UpdateAccountInDatabase() (err error) {
	err = db.GetGormDB().Table(library.TB_ACCOUNTS).Where("id = ?", ac.Account.Id).Updates(&ac.Account).Error
	return

	// c.JSON(http.StatusOK, a)
}

func (ac *AccoReferences) DeleteAccountInDatabase() (err error) {
	err = db.GetGormDB().Table(library.TB_ACCOUNTS).Where("cpf_cnpj = ?", ac.Account.CpfCnpj).Delete(&ac.Account).Error
	return

	// c.JSON(http.StatusOK, gin.H{"Account deleted": a})
}

// ----------------------------------------< Special case >---------------------------------------- \\

func DeleteByCpf_Cnpj(c *gin.Context, a *model.Account) {

	if err := db.GetGormDB().Table(library.TB_ACCOUNTS).Where("cpf_cnpj = ?", a.CpfCnpj).Delete(&a).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"Account deleted": a})
}

func GetAccountById(id int) (model.Account, error) {
	var NewAccount *model.Account
	if err := db.GetGormDB().Table(library.TB_ACCOUNTS).Where("id = ?", id).First(&NewAccount).Error; err != nil {
		return *NewAccount, err
	}
	return *NewAccount, nil
}
