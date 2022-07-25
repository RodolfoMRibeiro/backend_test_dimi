package predefinedData

import (
	entity_category "transaction/module/category/entity"
	entity_status "transaction/module/status/entity"
	entity_transaction "transaction/module/transaction/entity"
	entity_user "transaction/module/user/entity"

	"gorm.io/gorm"
)

func Load(db *gorm.DB) {
	categoryData(db)
	statusData(db)
}

func categoryData(db *gorm.DB) {
	categories := []entity_category.Category{
		{Id: 1, Name: "Lojista", Users: []entity_user.User{}},
		{Id: 2, Name: "Comum", Users: []entity_user.User{}},
	}

	for _, category := range categories {
		db.Table("tb_categories").Create(&category)
	}
}

func statusData(db *gorm.DB) {
	statusArr := []entity_status.Status{
		{Id: 1, Name: "Autorizado", Transaction: []entity_transaction.Transaction{}},
		{Id: 2, Name: "Não Autorizado", Transaction: []entity_transaction.Transaction{}},
	}

	for _, status := range statusArr {
		db.Table("tb_status").Create(&status)
	}
}
