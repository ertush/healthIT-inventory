package models

type Book struct {
	Id     int    `json:"id" gorm:"primaryKey"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Desc   string `json:"desc"`
}

type Stationery struct {
}

type Electronics struct {
}

type Furniture struct {
}
