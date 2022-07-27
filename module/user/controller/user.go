package controller

import (
	"net/http"
	"transaction/db"

	// controller_transaction "transaction/module/transaction/controller"
	"transaction/module/account/controller"

	entity_user "transaction/module/user/entity"
	"transaction/util"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {

	var newUser *entity_user.User = &entity_user.User{}

	if check(c, c.BindJSON(&newUser)) && checkEmailAndCPForCPNJ(newUser) {
		AddUserToDataBase(c, newUser)
	}
}

func FindUser(c *gin.Context) {

	var newUsers []entity_user.User = []entity_user.User{}

	if err := db.DB.Find(&newUsers).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	for i := 0; i < len(newUsers); i++ {
		newUsers[i].Account = controller.GetAccountsFromUser(newUsers[i].CpfCnpj)
	}

	c.JSON(http.StatusFound, newUsers)
}

func UploadUser(c *gin.Context) {

	var newUser *entity_user.User = &entity_user.User{}

	if check(c, c.BindJSON(&newUser)) && checkEmailAndCPForCPNJ(newUser) {
		UpdateUserInDataBase(c, newUser)
	}
}

func DeleteUser(c *gin.Context) {

	var newUser *entity_user.User = &entity_user.User{}

	if check(c, c.BindJSON(&newUser)) && checkEmailAndCPForCPNJ(newUser) {
		DeleteUserInDataBase(c, newUser)
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

func checkEmailAndCPForCPNJ(u *entity_user.User) (boolean bool) {

	cpfORcnpj, ok := util.VerifyingCPForCNPJ(util.LetOnlyNumbers(util.TrimAllSpacesInString(u.CpfCnpj)))
	if ok && util.IsEmailValid(util.TrimAllSpacesInString(u.Email)) {
		boolean = true
		u.CpfCnpj = cpfORcnpj
	}
	return
}

// -----------------------------------------< feed database >----------------------------------------- \\

func AddUserToDataBase(c *gin.Context, u *entity_user.User) {
	if err := db.DB.Table("tb_users").Create(&u).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusCreated, u)
}

func FindUserInDataBase(c *gin.Context, us *[]entity_user.User) {
	if err := db.DB.Find(us).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusFound, us)
}

func UpdateUserInDataBase(c *gin.Context, u *entity_user.User) {
	if err := db.DB.Table("tb_users").Where("cpf_cnpj = ?", u.CpfCnpj).Updates(&u).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, u)
}

func DeleteUserInDataBase(c *gin.Context, u *entity_user.User) {

	c.JSON(http.StatusInternalServerError, u)
}

func GetUserByAccountId(id int) (entity_user.User, error) {
	var newUser = &entity_user.User{}
	account, err := controller.GetAccountById(id)

	if err != nil {
		return *newUser, err
	}

	if err := db.DB.Table("tb_users").Where("cpf_cnpj = ?", account.CpfCnpj).First(newUser).Error; err != nil {
		return *newUser, err
	}

	return *newUser, nil
}
