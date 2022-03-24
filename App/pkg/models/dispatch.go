package models

import "time"

type Dispatch struct {
	Id         int       `json:"id" gorm:"primaryKey"`
	Name       string    `json:"name" gorm:"not null"`
	Facility   string    `json:"facility" gorm:"not null"`
	Equipment  string    `json:"equipment" gorm:"not null"`
	DateIssued time.Time `json:"date_issued" gorm:"not null"`
}
