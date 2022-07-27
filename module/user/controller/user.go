package controller

import (
	"net/http"
	"transaction/db"
	"transaction/library"

	"transaction/module/account/controller"
	entity_user "transaction/module/user/entity"
	"transaction/util"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var newUser *entity_user.User = &entity_user.User{}
	if containsError(c, c.BindJSON(&newUser)) && checkEmailAndCpf_Cnpf(newUser) {
		AddUserToDatabase(c, newUser)
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
	if containsError(c, c.BindJSON(&newUser)) && checkEmailAndCpf_Cnpf(newUser) {
		UpdateUserInDatabase(c, newUser)
	}
}

func DeleteUser(c *gin.Context) {
	c.IndentedJSON(http.StatusNotFound, "Sorry, but this method hasn't been developed yet")
}

// -------------------------------------------< Aux funcs >------------------------------------------- \\

func containsError(c *gin.Context, err error) bool {
	if err != nil {
		c.IndentedJSON(http.StatusNotAcceptable, "wrong data inserted") // 406
		return false
	}
	return true
}

func checkEmailAndCpf_Cnpf(u *entity_user.User) (boolean bool) {
	cpf_cnpj, ok := util.VerifyingCPForCNPJ(util.LetOnlyNumbers(util.TrimAllSpacesInString(u.CpfCnpj)))
	if ok && util.IsEmailValid(util.TrimAllSpacesInString(u.Email)) {
		boolean = true
		u.CpfCnpj = cpf_cnpj
	}
	return
}

// -----------------------------------------< feed database >----------------------------------------- \\

func AddUserToDatabase(c *gin.Context, u *entity_user.User) {
	if err := db.DB.Table(library.TB_USERS).Create(&u).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusCreated, u)
}

func FindUserInDatabase(c *gin.Context, us *[]entity_user.User) {
	if err := db.DB.Find(us).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusFound, us)
}

func UpdateUserInDatabase(c *gin.Context, u *entity_user.User) {
	if err := db.DB.Table(library.TB_USERS).Where("cpf_cnpj = ?", u.CpfCnpj).Updates(&u).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, u)
}

func GetUserByAccountId(id int) (entity_user.User, error) {
	var newUser = &entity_user.User{}
	account, err := controller.GetAccountById(id)

	if err != nil {
		return *newUser, err
	}

	if err := db.DB.Table(library.TB_USERS).Where("cpf_cnpj = ?", account.CpfCnpj).First(newUser).Error; err != nil {
		return *newUser, err
	}
	return *newUser, nil
}
