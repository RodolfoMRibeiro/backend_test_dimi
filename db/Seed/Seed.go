package seed

import (
	"transaction/library"
	entity_category "transaction/module/category/entity"
	entity_status "transaction/module/status/entity"
	entity_transaction "transaction/module/transaction/entity"
	entity_user "transaction/module/user/entity"

	"gorm.io/gorm"
)

func Handler(db *gorm.DB) {
	populateCategoryData(db)
	populateStatusData(db)
}

func populateCategoryData(db *gorm.DB) {
	categories := []entity_category.Category{
		{Id: 1, Name: "Lojista", Users: []entity_user.User{}},
		{Id: 2, Name: "Comum", Users: []entity_user.User{}},
	}

	for _, category := range categories {
		if err := db.Table(library.TB_CATEGORIES).Create(&category).Error; err != nil {
			break
		}
	}
}

func populateStatusData(db *gorm.DB) {
	statusArr := []entity_status.Status{
		{Id: 1, Name: "Autorizado", Transaction: []entity_transaction.Transaction{}},
		{Id: 2, Name: "Não Autorizado", Transaction: []entity_transaction.Transaction{}},
	}

	for _, status := range statusArr {
		if err := db.Table(library.TB_STATUS).Create(&status).Error; err != nil {
			break
		}
	}
}
