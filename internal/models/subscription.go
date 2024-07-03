package models

import (
	"strings"

	"gorm.io/gorm"
)

type Subscription struct {
	ID           int    `json:"id" gorm:"primaryKey"`
	Username     string `json:"username"`
	EmployeeName string `json:"employee_name"`
	ChatID       int64  `json:"chat_id"`
}

func (s *Subscription) BeforeSave(tx *gorm.DB) (err error) {
	s.EmployeeName = strings.ToLower(s.EmployeeName)
	return
}
