package main

import (
	"log"
	"net/http"
	"time"

	"birthday-notification-service/internal/bot"
	"birthday-notification-service/internal/db"
	"birthday-notification-service/internal/handlers"
	"birthday-notification-service/internal/middleware"
	"birthday-notification-service/internal/models"
	"birthday-notification-service/internal/services"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/robfig/cron/v3"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Ошибка загрузки файла .env")
	}

	db.Init()

	// Добавление тестовых данных (для примера)
	db.DB.Create(&models.Employee{Name: "райан гослинг", Birthday: models.Date{Time: time.Date(1980, time.November, 12, 0, 0, 0, 0, time.UTC)}})
	db.DB.Create(&models.Employee{Name: "дуэйн джонсон", Birthday: models.Date{Time: time.Date(1972, time.May, 2, 0, 0, 0, 0, time.UTC)}})
	db.DB.Create(&models.Employee{Name: "эмили блант", Birthday: models.Date{Time: time.Date(1983, time.February, 23, 0, 0, 0, 0, time.UTC)}})

	bot.Init()

	c := cron.New()
	c.AddFunc("@daily", func() { services.CheckAndSendBirthdayNotifications() })
	c.Start()

	r := mux.NewRouter()

	r.HandleFunc("/login", handlers.Login).Methods("POST")
	r.Handle("/employees", middleware.JWTAuthMiddleware(http.HandlerFunc(handlers.GetEmployees))).Methods("GET")
	r.Handle("/subscribe", middleware.JWTAuthMiddleware(http.HandlerFunc(handlers.Subscribe))).Methods("POST")
	r.Handle("/unsubscribe", middleware.JWTAuthMiddleware(http.HandlerFunc(handlers.Unsubscribe))).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", r))
}
