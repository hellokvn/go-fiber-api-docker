package models

type Book struct {
	Id          int    `json:"id" gorm:"primaryKey"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
}
