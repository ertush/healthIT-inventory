package models

type Electronics struct {
	Id        int              `json:"id" gorm:"primaryKey"`
	Equipment models.Equipment `json:"equipment" gorm:"foreignKeys"`
	Name      string           `json:"author"`
	Desc      string           `json:"desc"`
}
