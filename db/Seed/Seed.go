package seed

import (
	model "transaction/db/model"
	"transaction/library"

	"gorm.io/gorm"
)

func Handler(db *gorm.DB) {

	populateCategoryData(db)
	populateStatusData(db)
}

func populateCategoryData(db *gorm.DB) {
	categories := []model.Category{
		{Id: 1, Name: "Lojista", Users: []model.User{}},
		{Id: 2, Name: "Comum", Users: []model.User{}},
	}

	for _, category := range categories {
		if err := db.Table(library.TB_CATEGORIES).Find(&category).Error; err != nil {
			db.Table(library.TB_CATEGORIES).Create(category)
		}
	}
}

func populateStatusData(db *gorm.DB) {
	statusArr := []model.Status{
		{Id: 1, Name: "Autorizado", Transaction: []model.Transaction{}},
		{Id: 2, Name: "Não Autorizado", Transaction: []model.Transaction{}},
	}

	for _, status := range statusArr {
		if err := db.Table(library.TB_STATUS).Find(&status).Error; err != nil {
			db.Table(library.TB_STATUS).Create(&status)
		}
	}
}
