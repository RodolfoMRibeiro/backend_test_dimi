package db

import (
	"fmt"
	entity_env_vars "transaction/configs/entity"
	entity_account "transaction/module/account/entity"
	entity_category "transaction/module/category/entity"
	entity_status "transaction/module/status/entity"
	entity_transaction "transaction/module/transaction/entity"
	entity_user "transaction/module/user/entity"
	"transaction/util"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

var envVars entity_env_vars.EnvironmentVariables

func Load() {
	godotenv.Load(".env")
	envVars.FeedStruct()
	connectDatabase()
}

func connectDatabase() {
	databaseStringConfig := fmt.Sprintf(

		"%s:%s@tcp(%s:%s)/%s%s",

		envVars.User,
		envVars.Password,
		envVars.Host,
		envVars.Port,
		envVars.Database,
		envVars.Configs,
	)

	database, err := gorm.Open(mysql.Open(databaseStringConfig), &gorm.Config{})

	util.PresentatePanicErros(err)

	migrate(database)

	db = database
}

func migrate(db *gorm.DB) {
	db.AutoMigrate(&entity_account.Account{})
	db.AutoMigrate(&entity_category.Category{})
	db.AutoMigrate(&entity_status.Status{})
	db.AutoMigrate(&entity_transaction.Transaction{})
	db.AutoMigrate(&entity_user.User{})

	// db.Table("tb_accounts").AutoMigrate(&entity_account.Account{})
	// db.Table("tb_categories").AutoMigrate(&entity_category.Category{})
	// db.Table("tb_status").AutoMigrate(&entity_status.Status{})
	// db.Table("tb_transactions").AutoMigrate(&entity_transaction.Transaction{})
	// db.Table("tb_users").AutoMigrate(&entity_user.User{})
}

func GetGormDB() *gorm.DB {
	return db
}
