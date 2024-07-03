package db

import (
	"birthday-notification-service/internal/models"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Init инициализирует базу данных
func Init() {
	var err error
	DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Не удалось подключиться к базе данных: %v", err)
	}

	err = DB.AutoMigrate(&models.Employee{}, &models.Subscription{})
	if err != nil {
		log.Fatalf("Не удалось выполнить миграцию схемы базы данных: %v", err)
	}
}
