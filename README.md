# Birthday Notification Service

Этот проект представляет собой сервис для удобного поздравления сотрудников с днем рождения. Сервис включает в себя API для управления сотрудниками и подписками на уведомления о днях рождения, а также интеграцию с Telegram ботом.

## Структура проекта

```plaintext
birthday-notification-service/
├── cmd/
│   └── main.go                    # Главный файл запуска приложения
├── internal/
│   ├── bot/
│   │   └── bot.go                 # Логика Telegram бота
│   ├── db/
│   │   └── db.go                  # Инициализация базы данных
│   ├── handlers/
│   │   ├── auth_handler.go        # Обработчики для авторизации
│   │   └── subscription_handler.go # Обработчики для подписки и отписки
│   ├── middleware/
│   │   └── auth_middleware.go     # Middleware для проверки JWT токенов
│   ├── models/
│   │   ├── employee.go            # Модели данных для сотрудников
│   │   └── subscription.go        # Модели данных для подписок
│   └── services/
│       └── notification_service.go # Логика отправки уведомлений
├── test/
│   └── subscription_handler_test.go # Тесты для подписки и отписки
├── go.mod                         # Go module file
└── go.sum                         # Go sum file
```
## Запуск приложения

### Требования

- Go 1.22.4 или выше
- Установленный SQLite3
- Переменная окружения `TELEGRAM_BOT_TOKEN` с токеном вашего Telegram бота

### Шаги для запуска

1. Клонируйте репозиторий:

   ```sh
   git clone git@github.com:gratefultolord/birthday-notification-service.git
   cd birthday-notification-service
2. Установите зависимости:
   ```sh
   go mod download
3. Создайте файл .env в корне проекта и добавьте переменные окружения:
   ```sh
   TELEGRAM_BOT_TOKEN=your_telegram_bot_token
4. Запустите приложение:
   ```sh
   go run cmd/main.go

## Запуск тестов
```sh
go test ./test/...


