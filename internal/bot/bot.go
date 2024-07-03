package bot

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"birthday-notification-service/internal/db"
	"birthday-notification-service/internal/models"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var Bot *tgbotapi.BotAPI

func Init() {
	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	if botToken == "" {
		log.Fatal("Переменная окружения TELEGRAM_BOT_TOKEN обязательна")
	}

	var err error
	Bot, err = tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Fatalf("Не удалось создать бота: %v", err)
	}

	Bot.Debug = true
	log.Printf("Авторизован под аккаунтом %s", Bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := Bot.GetUpdatesChan(u)

	go func() {
		for update := range updates {
			if update.Message == nil {
				continue
			}

			switch update.Message.Command() {
			case "start":
				handleStartCommand(update.Message)
			case "subscribe":
				handleSubscribeCommand(update.Message)
			case "unsubscribe":
				handleUnsubscribeCommand(update.Message)
			default:
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Неизвестная команда")
				Bot.Send(msg)
			}
		}
	}()
}

func handleStartCommand(message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Добро пожаловать в бота уведомлений о днях рождения! Используйте /subscribe <имя_фамилия_сотрудника> для подписки и /unsubscribe <имя_фамилия_сотрудника> для отписки.")
	Bot.Send(msg)
}

func handleSubscribeCommand(message *tgbotapi.Message) {
	args := message.CommandArguments()
	if args == "" {
		msg := tgbotapi.NewMessage(message.Chat.ID, "Пожалуйста, укажите имя и фамилию сотрудника.")
		Bot.Send(msg)
		return
	}

	employeeName := strings.ToLower(strings.TrimSpace(args))
	username := message.Chat.UserName

	// Найти сотрудника по имени
	var employee models.Employee
	db.DB.Find(&employee, "LOWER(name) = ?", employeeName)
	if employee.ID == 0 {
		msg := tgbotapi.NewMessage(message.Chat.ID, "Сотрудник не найден.")
		Bot.Send(msg)
		return
	}

	subscription := models.Subscription{
		Username:     username,
		EmployeeName: employeeName,
	}

	db.DB.Create(&subscription)

	// Вычисление количества дней до дня рождения
	today := time.Now()
	nextBirthday := time.Date(today.Year(), employee.Birthday.Month(), employee.Birthday.Day(), 0, 0, 0, 0, today.Location())
	if today.After(nextBirthday) {
		nextBirthday = nextBirthday.AddDate(1, 0, 0)
	}
	daysUntilBirthday := nextBirthday.Sub(today).Hours() / 24

	var notificationMessage string
	if daysUntilBirthday > 30 {
		months := int(daysUntilBirthday / 30)
		days := int(daysUntilBirthday) % 30
		notificationMessage = fmt.Sprintf("Вы успешно подписались на уведомления о сотруднике %s. До его дня рождения осталось %d месяцев и %d дней.", args, months, days)
	} else {
		notificationMessage = fmt.Sprintf("Вы успешно подписались на уведомления о сотруднике %s. До его дня рождения осталось %.0f дней.", args, daysUntilBirthday)
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, notificationMessage)
	Bot.Send(msg)
}

func handleUnsubscribeCommand(message *tgbotapi.Message) {
	args := message.CommandArguments()
	if args == "" {
		msg := tgbotapi.NewMessage(message.Chat.ID, "Пожалуйста, укажите имя и фамилию сотрудника.")
		Bot.Send(msg)
		return
	}

	employeeName := strings.ToLower(strings.TrimSpace(args))
	username := message.Chat.UserName

	db.DB.Where("username = ? AND employee_name = ?", username, employeeName).Delete(&models.Subscription{})

	msg := tgbotapi.NewMessage(message.Chat.ID, "Вы успешно отписались от уведомлений о сотруднике "+args+".")
	Bot.Send(msg)
}
