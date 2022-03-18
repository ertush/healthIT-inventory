package models

type User struct {
	Id       uint   `json:"id" gorm:"primaryKey not null auto increment"`
	Name     string `json:"name" gorm:"not null"`
	Email    string `json:"email" gorm:"unique not null"`
	Password []byte `json:"-" gorm:"not null"`
	Role     string `json:"role"`
}
