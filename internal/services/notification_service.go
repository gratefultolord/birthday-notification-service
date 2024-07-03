package services

import (
	"birthday-notification-service/internal/bot"
	"birthday-notification-service/internal/db"
	"birthday-notification-service/internal/models"
	"fmt"
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func SendBirthdayNotification(username string, employee models.Employee) {
	chatID, err := getChatIDByUsername(username)
	if err != nil {
		fmt.Printf("Не удалось найти пользователя %s: %v\n", username, err)
		return
	}

	message := fmt.Sprintf("Сегодня день рождения у сотрудника %s!", employee.Name)
	msg := tgbotapi.NewMessage(chatID, message)
	bot.Bot.Send(msg)
}

func getChatIDByUsername(username string) (int64, error) {
	var subscription models.Subscription
	if err := db.DB.Where("username = ?", username).First(&subscription).Error; err != nil {
		return 0, err
	}
	return subscription.ChatID, nil
}

// CheckAndSendBirthdayNotifications проверяет дни рождения и отправляет уведомления подписанным пользователям
func CheckAndSendBirthdayNotifications() {
	var employees []models.Employee
	db.DB.Find(&employees)

	today := time.Now().Format("2006-01-02")

	for _, employee := range employees {
		if employee.Birthday.Format("2006-01-02") == today {
			var subscriptions []models.Subscription
			db.DB.Where("employee_name = ?", strings.ToLower(employee.Name)).Find(&subscriptions)
			for _, subscription := range subscriptions {
				SendBirthdayNotification(subscription.Username, employee)
			}
		}
	}
}
