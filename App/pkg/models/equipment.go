package models

type Equipment struct {
	Id         int    `json:"id" gorm:"primaryKey"`
	Name       string `json:"name"`
	Desc       string `json:"desc"`
	BatchNo    string `json:"batchNo"`
	CategoryId string `json:"categoryid gorm:"foreignKey" reference:id`
}
