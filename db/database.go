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

var DB *gorm.DB
var envVars entity_env_vars.EnvironmentVariables

func LoadEnv() {
	godotenv.Load(".env")
	envVars.FeedStruct()
}

func ConnectDatabase() {
	databaseStringConfig := fmt.Sprintf(

		"%s:%s@tcp(%s:%s)/%s%s",

		envVars.User,
		envVars.Password,
		envVars.Host,
		envVars.Port,
		envVars.Database,
		envVars.Configs,
	)

	db, err := gorm.Open(mysql.Open(databaseStringConfig), &gorm.Config{})

	util.PresentateErros(err)

	loadMigrations(db)

	DB = db
}

func loadMigrations(db *gorm.DB) {
	db.AutoMigrate(&entity_account.Account{})
	db.AutoMigrate(&entity_category.Category{})
	db.AutoMigrate(&entity_status.Status{})
	db.AutoMigrate(&entity_transaction.Transaction{})
	db.AutoMigrate(&entity_user.User{})
}
