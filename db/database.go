package db

import (
	"fmt"
	vars "transaction/db/entity"
	account "transaction/module/account/entity"
	category "transaction/module/category/entity"
	status "transaction/module/status/entity"
	transaction "transaction/module/transaction/entity"
	user "transaction/module/user/entity"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var envVars vars.EnvironmentVariables

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

	db, err := gorm.Open(mysql.Open(databaseStringConfig), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	migrate(db)

	DB = db
}

func migrate(db *gorm.DB) {
	db.AutoMigrate(&account.Account{})
	db.AutoMigrate(&category.Category{})
	db.AutoMigrate(&status.Status{})
	db.AutoMigrate(&transaction.Transaction{})
	db.AutoMigrate(&user.User{})
}
