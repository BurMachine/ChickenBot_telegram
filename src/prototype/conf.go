package main

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

const (
	TELEGRAM_BOT_API_KEY       = "5775513785:AAGy6Ht6IYgaZUVfLOmyyYiviwtJfJhmKu8" // API  ключ, который мы получили у BotFather
	POSTGRESS_CONNECTION_URL   = "localhost"                                      // Адрес сервера PostgressDB
	POSTGRESS_DATABASE_NAME    = "regbot"                                         // Название базы данных
	POSTGRESS_COLLECTION_USERS = "users"                                          // Название таблицы
)

type User struct {
	Chat_ID int64
}

type TelegramBot struct {
	API                   *tgbotapi.BotAPI        // API телеграмма
	Updates               tgbotapi.UpdatesChannel // Канал обновлений
	ActiveContactRequests []int64                 // ID чатов, от которых мы ожидаем номер
}

type user struct {
	state  int // 0, 1, 2, 3
	name   string
	login  string
	campus string
	role   int
}

<<<<<<< HEAD
var signMap map[int64]*user
=======
type events struct {
	eType       string
	description string
	uniqueCode  string
	startTime   string //проверить типб в БД timestamp
	expiresTime string //проверить типб в БД timestamp
}

var signMap map[int]*user
>>>>>>> b34a0062c56fdf46a9bca8eed437124886333504

var StartMenuKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Регистрация"),
	),
)

var CampusMenuKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Казань"),
		tgbotapi.NewKeyboardButton("Москва"),
		tgbotapi.NewKeyboardButton("Новосибирск"),
	),
)
