package handlers

import (
	"birthday-notification-service/internal/db"
	"birthday-notification-service/internal/models"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/context"
)

type SubscriptionRequest struct {
	EmployeeName string `json:"employee_name"`
}

// Subscribe добавляет подписку на уведомления о дне рождения
func Subscribe(w http.ResponseWriter, r *http.Request) {
	var req SubscriptionRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Некорректный запрос", http.StatusBadRequest)
		return
	}

	username := context.Get(r, "username").(string)

	subscription := models.Subscription{
		Username:     username,
		EmployeeName: strings.ToLower(req.EmployeeName),
	}

	db.DB.Create(&subscription)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(subscription)
}

// Unsubscribe удаляет подписку на уведомления о дне рождения
func Unsubscribe(w http.ResponseWriter, r *http.Request) {
	var req SubscriptionRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Некорректный запрос", http.StatusBadRequest)
		return
	}

	username := context.Get(r, "username").(string)

	db.DB.Where("username = ? AND employee_name = ?", username, strings.ToLower(req.EmployeeName)).Delete(&models.Subscription{})

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Успешная отписка"})
}
