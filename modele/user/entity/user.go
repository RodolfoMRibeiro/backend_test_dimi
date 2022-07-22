package entity

type User struct {
	Cpf_Cnpj    string      `gorm:"primaryKey;autoIncrement:false"`
	Id_category Id_category `gorm:"foreignKey:Id"`
	Fullname    string
	Password    string
}
