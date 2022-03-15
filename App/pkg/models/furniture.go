package models

// import "github.com/ertush/healthIT-inventory/App/pkg/models"

type Furniture struct {
	Id        int       `json:"id" gorm:"primaryKey"`
	Equipment Equipment `json:"equipment" gorm:"foreignKeys"`
}
