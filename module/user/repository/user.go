package repository

import (
	// "fmt"
	"transaction/db"
	entity_user "transaction/module/user/entity"
)

func FindUserById(id int, u *entity_user.User) {
	db.DB.Where("id = ?", id).First(u)
}

func AddUserToDataBase(u *entity_user.User) {
	db.DB.Create(&u)
}
