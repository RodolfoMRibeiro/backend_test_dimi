package repository

import (
	"transaction/db"

	"gorm.io/gorm"
)

type RepoReferences interface {
	UserReferences
	AccoReferences
	TranReferences
}

type DbReferences struct {
	DB             *gorm.DB
	RepoReferences interface{}
}

func (Db *DbReferences) GetDataBase() {
	Db.DB = db.GetGormDB()
}
