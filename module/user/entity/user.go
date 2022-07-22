package entity

type User struct {
	Cpf_Cnpj    string `gorm:"primaryKey;autoIncrement:false"`
	Id_category int    `gorm:"foreignKey:Id"`
	Fullname    string
	Password    string
}
