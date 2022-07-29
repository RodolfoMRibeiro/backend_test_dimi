package controller

import (
	"net/http"
	"transaction/library"
	model "transaction/module/models"

	repository "transaction/module/repositories"
	"transaction/util"

	"github.com/gin-gonic/gin"
)

func FindTransaction(c *gin.Context) {
	var new *repository.TranReferences

	if !util.ContainsError(c.BindJSON(&new.Transactions)) {
		err := new.FindTransactionsInDatabase(c)
		util.FoundOrNotStatusReturn(err, c, new.Transactions)
	}
}

func CreateTransaction(c *gin.Context) {
	var NewTransaction *model.Transaction

	if util.ContainsError(c.BindJSON(&NewTransaction)) {
		c.IndentedJSON(http.StatusNotAcceptable, "wrong data inserted") // 406
		return
	}

	NewTransaction.ValidateTransaction()
	if NewTransaction.IdPayer != NewTransaction.IdPayee && NewTransaction.IdStatus == library.STORE_KEEPER_STATUS && !isLojista(NewTransaction.IdPayer) {
		if err := repository.BeginTransaction(NewTransaction); err != nil {
			c.IndentedJSON(http.StatusInternalServerError, err)
			return
		}
		c.JSON(http.StatusOK, gin.H{"New user registred": NewTransaction})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Ops! something went wrong"})
	}
}

// ----------------------------------------< Aux func >---------------------------------------- \\

func isLojista(AccountId int) bool {
	user, _ := repository.GetUserByAccountId(AccountId)
	if user.IdCategory == 1 {
		return true
	} else {
		return false
	}
}
