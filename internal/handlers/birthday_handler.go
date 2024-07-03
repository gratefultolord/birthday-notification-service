package handlers

import (
	"birthday-notification-service/internal/models"
	"encoding/json"
	"net/http"
	"time"
)

var employees = []models.Employee{
	{ID: 1, Name: "John Doe", Birthday: parseDate("1980-01-01")},
	{ID: 2, Name: "Jane Smith", Birthday: parseDate("1990-02-15")},
	{ID: 3, Name: "Emily Johnson", Birthday: parseDate("1985-07-30")},
}

func parseDate(dateStr string) models.Date {
	date, _ := time.Parse("2006-01-02", dateStr)
	return models.Date{Time: date}
}

func GetEmployees(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(employees)
}
