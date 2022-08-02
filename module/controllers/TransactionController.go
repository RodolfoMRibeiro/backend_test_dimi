package controller

import (
	"net/http"
	"transaction/library"

	repository "transaction/module/repositories"
	service "transaction/module/services"
	"transaction/util"

	"github.com/gin-gonic/gin"
)

func FindTransaction(c *gin.Context) {
	var new repository.TranReferences

	err := new.FindTransactionsInDatabase()
	service.FoundOrNotStatusReturn(err, c, new.Transactions)
}

func CreateTransaction(c *gin.Context) {
	var new repository.TranReferences

	if util.ContainsError(c.BindJSON(&new.Transaction)) {
		c.IndentedJSON(http.StatusNotAcceptable, "wrong data inserted") // 406
		return
	}

	new.Transaction.ValidateTransaction()
	if new.Transaction.IdPayer != new.Transaction.IdPayee && new.Transaction.IdStatus == library.STORE_KEEPER_STATUS && !isStoreKeeper(new.Transaction.IdPayer) {
		if err := repository.BeginTransaction(new.Transaction); err != nil {
			c.IndentedJSON(http.StatusInternalServerError, err)
			return
		}
		c.JSON(http.StatusOK, gin.H{"New user registred": new.Transaction})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Ops! something went wrong"})
	}
}

// ----------------------------------------< Aux func >---------------------------------------- \\

func isStoreKeeper(AccountId int) bool {
	user, _ := repository.GetUserByAccountId(AccountId)
	if user.IdCategory == 1 {
		return true
	} else {
		return false
	}
}
