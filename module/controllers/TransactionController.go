package controller

import (
	"net/http"
	"transaction/library"

	model "transaction/db/model"
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

	new.ValidateTransaction()
	if isValidPayer(new.Transaction) {
		if err := repository.BeginTransaction(new.Transaction); err != nil {
			c.IndentedJSON(http.StatusInternalServerError, err)
			return
		}
		c.JSON(http.StatusOK, gin.H{"New user registered": new.Transaction})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Ops! something went wrong"})
	}
}

// ----------------------------------------< Aux func >---------------------------------------- \\

func isValidPayer(transac *model.Transaction) bool {
	payerIsNotPayee := transac.IdPayer != transac.IdPayee
	isStoreKeeper := transac.IdStatus == library.STORE_KEEPER_STATUS

	return payerIsNotPayee && isStoreKeeper
}
