package entity

import (
	"transaction/module/user/entity"
)

type Category struct {
	Id    int           `gorm:"primaryKey"`
	Name  string        `gorm:"type:varchar(15)"`
	Users []entity.User `gorm:"foreignKey:IdCategory"`
}

type Tabler interface {
	TableName() string
	FeedAssociatedTable()
}

func (Category) TableName() string {
	return "tb_categories"
}

// func (Category) FeedAssociatedTable() {

// 	var c = []Category{{Id: 1, Name: "Accepted"}, {Id: 2, Name: "Denied"}}
// 	db.DB.Table("tb_users").Create(c)
// }
