package models

type Medicine struct {
	Id        int       `json:"id" gorm:"primaryKey"`
	Equipment Equipment `json:"equipment" gorm:"foreignKeys"`
	Name      string    `json:"author"`
	Desc      string    `json:"desc"`
}