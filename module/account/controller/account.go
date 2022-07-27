package controller

import (
	"net/http"
	"transaction/db"
	entity_account "transaction/module/account/entity"
	"transaction/util"

	"github.com/gin-gonic/gin"
)

func CreateAccount(c *gin.Context) {

	var NewAccount *entity_account.Account = &entity_account.Account{}

	if check(c, c.BindJSON(NewAccount)) {
		AddAccountToDataBase(c, NewAccount)
	}

	c.JSON(http.StatusOK, gin.H{"New transaction registred": NewAccount})
}

func FindAccount(c *gin.Context) {

	var NewAccounts *[]entity_account.Account = &[]entity_account.Account{}

	FindAccountInDataBase(c, NewAccounts)
}

func UpdateAccount(c *gin.Context) {

	var NewAccount *entity_account.Account = &entity_account.Account{}

	if check(c, c.BindJSON(NewAccount)) {
		UpdateAccountInDataBase(c, NewAccount)
	}
}

func DeleteAccount(c *gin.Context) {

	var NewAccount *entity_account.Account = &entity_account.Account{}

	if check(c, c.BindJSON(NewAccount)) {
		DeleteAccountInDataBase(c, NewAccount)
	}
}

func DeleteAccountsByCpf_Cnpj(c *gin.Context, cpf_cnpj string) {

	var NewAccount *entity_account.Account = &entity_account.Account{}
	NewAccount.CpfCnpj = cpf_cnpj

	if checkCPForCPNJ(NewAccount) && check(c, c.BindJSON(NewAccount)) {
		DeleteEverything(c, NewAccount)
	}
}

// -------------------------------------------< Aux funcs >------------------------------------------- \\

func check(c *gin.Context, err error) bool {

	if err != nil {
		c.IndentedJSON(http.StatusNotAcceptable, "wrong data inserted") // 406
		return false
	}
	return true
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

	if err := db.DB.Table("tb_accounts").Where("cpf_cnpj = ?", cpf_cnpj).Find(&NewAccounts).Error; err != nil {
		return nil
	}

	return NewAccounts
}

// -----------------------------------------< feed database >----------------------------------------- \\

func AddAccountToDataBase(c *gin.Context, a *entity_account.Account) {

	if err := db.DB.Table("tb_accounts").Create(&a).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusCreated, gin.H{"New Account registred": a})
}

func FindAccountInDataBase(c *gin.Context, as *[]entity_account.Account) {

	if err := db.DB.Table("tb_accounts").Find(as).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusFound, as)
}

func UpdateAccountInDataBase(c *gin.Context, a *entity_account.Account) {

	if err := db.DB.Table("tb_accounts").Where("id = ?", a.Id).Updates(a).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, a)
}

func DeleteAccountInDataBase(c *gin.Context, a *entity_account.Account) {

	if err := db.DB.Table("tb_accounts").Where("id = ?", a.Id).Delete(a).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"Account deleted": a})
}

func DeleteEverything(c *gin.Context, a *entity_account.Account) {

	if err := db.DB.Table("tb_accounts").Where("cpf_cnpj = ?", a.CpfCnpj).Delete(a).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"Account deleted": a})
}

func GetAccountById(id int) (entity_account.Account, error) {
	var NewAccount *entity_account.Account = &entity_account.Account{}

	if err := db.DB.Table("tb_accounts").Where("id = ?", id).First(NewAccount).Error; err != nil {
		return *NewAccount, err
	}
	return *NewAccount, nil
}
