package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"birthday-notification-service/internal/db"
	"birthday-notification-service/internal/handlers"
	"birthday-notification-service/internal/models"

	"github.com/gorilla/context"
	"github.com/stretchr/testify/assert"
)

func TestSubscribe(t *testing.T) {
	db.Init()
	db.DB.Create(&models.Employee{Name: "тест сотрудник", Birthday: models.Date{Time: time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC)}})

	body := map[string]string{"employee_name": "тест сотрудник"}
	jsonBody, _ := json.Marshal(body)
	req, err := http.NewRequest("POST", "/subscribe", bytes.NewBuffer(jsonBody))
	assert.NoError(t, err)

	context.Set(req, "username", "тест пользователь")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.Subscribe)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var subscription models.Subscription
	result := db.DB.Where("username = ? AND employee_name = ?", "тест пользователь", "тест сотрудник").First(&subscription)
	assert.NoError(t, result.Error)
	assert.Equal(t, "тест пользователь", subscription.Username)
	assert.Equal(t, "тест сотрудник", subscription.EmployeeName)
}

func TestUnsubscribe(t *testing.T) {
	db.Init()
	db.DB.Create(&models.Employee{Name: "тест сотрудник", Birthday: models.Date{Time: time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC)}})
	db.DB.Create(&models.Subscription{Username: "тест пользователь", EmployeeName: "тест сотрудник"})

	body := map[string]string{"employee_name": "тест сотрудник"}
	jsonBody, _ := json.Marshal(body)
	req, err := http.NewRequest("POST", "/unsubscribe", bytes.NewBuffer(jsonBody))
	assert.NoError(t, err)

	context.Set(req, "username", "тест пользователь")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.Unsubscribe)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var subscription models.Subscription
	result := db.DB.Where("username = ? AND employee_name = ?", "тест пользователь", "тест сотрудник").First(&subscription)
	assert.Error(t, result.Error)
	assert.Equal(t, "record not found", result.Error.Error())
}
