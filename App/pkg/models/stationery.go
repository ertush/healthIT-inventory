package models

type Stationery struct {
	Id        int       `json:"id" gorm:"primaryKey"`
	Equipment Equipment `json:"equipment" gorm:"foreignKeys"`
	Name      string    `json:"author"`
	Desc      string    `json:"desc"`
}
