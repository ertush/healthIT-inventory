package models

type Categories struct {
	Id          int    `json:"id gorm:unique" gorm:"primaryKey notnull"`
	Name        string `json:"name" gorm:"not null"`
	Description string `json:"description" gorm:"not null"`
}
