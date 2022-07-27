package controller

import (
	"net/http"
	"transaction/db"
	"transaction/library"
	entity_account "transaction/module/account/entity"
	"transaction/util"

	"github.com/gin-gonic/gin"
)

func CreateAccount(c *gin.Context) {
	var NewAccount *entity_account.Account = &entity_account.Account{}
	if !containsError(c, c.BindJSON(NewAccount)) {
		AddAccountToDatabase(c, NewAccount)
	}
	c.JSON(http.StatusOK, gin.H{"New transaction registred": NewAccount})
}

func FindAccount(c *gin.Context) {
	var NewAccounts *[]entity_account.Account = &[]entity_account.Account{}
	FindAccountInDatabase(c, NewAccounts)
}

func UpdateAccount(c *gin.Context) {
	var NewAccount *entity_account.Account = &entity_account.Account{}
	if !containsError(c, c.BindJSON(NewAccount)) {
		UpdateAccountInDatabase(c, NewAccount)
	}
}

func DeleteAccount(c *gin.Context) {
	var NewAccount *entity_account.Account = &entity_account.Account{}
	if !containsError(c, c.BindJSON(NewAccount)) {
		DeleteAccountInDatabase(c, NewAccount)
	}
}

func DeleteAccountsByCpf_Cnpj(c *gin.Context, cpf_cnpj string) {
	var NewAccount *entity_account.Account = &entity_account.Account{}
	NewAccount.CpfCnpj = cpf_cnpj
	if checkCPForCPNJ(NewAccount) && !containsError(c, c.BindJSON(NewAccount)) {
		DeleteByCpf_Cnpj(c, NewAccount)
	}
}

// -------------------------------------------< Aux funcs >------------------------------------------- \\

func containsError(c *gin.Context, err error) bool {
	if err != nil {
		c.IndentedJSON(http.StatusNotAcceptable, "wrong data inserted") // 406
		return true
	}
	return false
}

func checkCPForCPNJ(a *entity_account.Account) (boolean bool) {
	if cpfORcnpj, ok := util.VerifyingCPForCNPJ(util.LetOnlyNumbers(util.TrimAllSpacesInString(a.CpfCnpj))); ok {
		boolean = true
		a.CpfCnpj = cpfORcnpj
	}
	return
}

func GetAccountsFromUser(cpf_cnpj string) []entity_account.Account {
	var NewAccounts []entity_account.Account = []entity_account.Account{}
	if err := db.DB.Table(library.TB_ACCOUNTS).Where("cpf_cnpj = ?", cpf_cnpj).Find(&NewAccounts).Error; err != nil {
		return nil
	}
	return NewAccounts
}

// -----------------------------------------< feed database >----------------------------------------- \\

func AddAccountToDatabase(c *gin.Context, a *entity_account.Account) {
	if err := db.DB.Table(library.TB_ACCOUNTS).Create(&a).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusCreated, gin.H{"New Account registred": a})
}

func FindAccountInDatabase(c *gin.Context, as *[]entity_account.Account) {
	if err := db.DB.Table(library.TB_ACCOUNTS).Find(as).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusFound, as)
}

func UpdateAccountInDatabase(c *gin.Context, a *entity_account.Account) {
	if err := db.DB.Table(library.TB_ACCOUNTS).Where("id = ?", a.Id).Updates(a).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, a)
}

func DeleteAccountInDatabase(c *gin.Context, a *entity_account.Account) {
	if err := db.DB.Table(library.TB_ACCOUNTS).Where("cpf_cnpj = ?", a.CpfCnpj).Delete(a).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"Account deleted": a})
}

func DeleteByCpf_Cnpj(c *gin.Context, a *entity_account.Account) {

	if err := db.DB.Table(library.TB_ACCOUNTS).Where("cpf_cnpj = ?", a.CpfCnpj).Delete(&a).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"Account deleted": a})
}

func GetAccountById(id int) (entity_account.Account, error) {
	var NewAccount *entity_account.Account = &entity_account.Account{}
	if err := db.DB.Table(library.TB_ACCOUNTS).Where("id = ?", id).First(NewAccount).Error; err != nil {
		return *NewAccount, err
	}
	return *NewAccount, nil
}
