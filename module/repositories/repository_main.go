package repository

import (
	"transaction/db"

	"gorm.io/gorm"
)

type DbReferences struct {
	DB             *gorm.DB
	RepoReferences interface{}
}

type RepoReferences interface {
	UserReferences
	AccoReferences
	TranReferences
}

func (Db *DbReferences) GetDataBase() {
	Db.DB = db.GetGormDB()
}
