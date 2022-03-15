package models

type EquipmentDetails struct {
	Id       int    `json:"id" gorm:"primaryKey"`
	Name     string `json:"name"`
	SerialNo string `json:"serialNo"`
	BatchNo  string `json: batchNo`
	Desc     string `json:"desc"`
}
