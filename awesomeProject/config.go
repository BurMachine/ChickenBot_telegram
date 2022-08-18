package main

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

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

type user struct {
	state  int // 0, 1, 2, 3
	name   string
	login  string
	campus string
	role   int
}
