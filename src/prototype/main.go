package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

var m map[int64]int

func main() {
	bot, err := tgbotapi.NewBotAPI(TELEGRAM_BOT_API_KEY) // подключаемся к боту с помощью токена
	if err != nil {
		log.Panic(err)
	}
	db := openDatabase()
	//bot.Debug = true
	//db, err := sql.Open("postgres", dbInfo)

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)
	flag := 0
	flag1 := 1
	m = make(map[int64]int)
	signMap = make(map[int64]*user)
	createMap = make(map[int64]*events)
	i := 0 // флаг регистрации(4 если все ок)
	for update := range updates {
		var msg tgbotapi.MessageConfig
		if update.Message != nil {
			if update.Message.IsCommand() {
				commandProc(update, bot, db, msg, &flag1, &flag)
			} else {
				nonCommandProc(update, bot, db, msg, &flag1, &flag, &i)
			}
		} else if update.CallbackQuery != nil {
			CallBackQueryProc(bot, update, msg, &flag1, &flag)
		}
	}
}
