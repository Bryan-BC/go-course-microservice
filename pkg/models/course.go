package models

type Course struct {
	Id          string `json:"id" gorm:"primary_key"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
