package repository

import (
	"fmt"
	"transaction/db"
	entity_user "transaction/module/user/entity"
)

func FindById(id int, u *entity_user.User) {
	db.DB.Where("id = ?", id).First(u)
}

func AddToDataBase(u *entity_user.User) {
	db.DB.Create(&u)
	fmt.Println("fisdjfiodiasofiodsofoijsoiafjoioi", u)
}
