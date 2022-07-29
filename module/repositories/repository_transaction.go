package repository

import (
	model "transaction/module/models"
)

type TranReferences struct {
	Transaction  *model.Transaction
	transactions *[]model.Transaction
}
