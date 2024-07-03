package models

type Employee struct {
	ID       int    `json:"id" gorm:"primaryKey"`
	Name     string `json:"name"`
	Birthday Date   `json:"birthday"`
}
