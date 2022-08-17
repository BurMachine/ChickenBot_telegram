package main

import tgbotapi "github.com/Syfaro/telegram-bot-api"

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

var signMap map[int]*user

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