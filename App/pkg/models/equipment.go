package models

type Equipment struct {
	Id          int    `json:"id" gorm:"primaryKey unique "`
	SerialNo    string `json:"serial_number" gorm:"not null"`
	Name        string `json:"name" gorm:"not null"`
	Description string `json:"description" gorm:"not null"`
	Category    string `json:"category" gorm:"ForeignKey:Name not null"`
	Status      string `json:"status"`
}
