package db

import (
	"fmt"
	config "transaction/configs"
	entity_account "transaction/module/account/entity"
	entity_category "transaction/module/category/entity"
	entity_status "transaction/module/status/entity"
	entity_transaction "transaction/module/transaction/entity"
	entity_user "transaction/module/user/entity"
	"transaction/util"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Load() {
	connectDatabase()
}

func connectDatabase() {
	databaseStringConfig := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s%s",

		config.Mysql.USER,
		config.Mysql.PASSWORD,
		config.Mysql.HOST,
		config.Mysql.PORT,
		config.Mysql.DB,
		config.Mysql.ADDITIONAL_CONFIGS,
	)

	database, err := gorm.Open(mysql.Open(databaseStringConfig), &gorm.Config{})

	util.PresentateErros(err)

	loadMigrations(database)

	db = database
}

func loadMigrations(db *gorm.DB) {
	db.AutoMigrate(&entity_account.Account{})
	db.AutoMigrate(&entity_category.Category{})
	db.AutoMigrate(&entity_status.Status{})
	db.AutoMigrate(&entity_transaction.Transaction{})
	db.AutoMigrate(&entity_user.User{})
}

func GetGormDB() *gorm.DB {
	return db
}
