package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

var m map[int64]int

func main() {
	bot, err := tgbotapi.NewBotAPI("5775513785:AAGy6Ht6IYgaZUVfLOmyyYiviwtJfJhmKu8") // подключаемся к боту с помощью токена
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
	m = make(map[int64]int)
	signMap = make(map[int64]*user)
	i := 0 // флаг регистрации(4 если все ок)
	regFlag := 0
	for update := range updates {
		var msg tgbotapi.MessageConfig
		if update.Message != nil {
			_, ok := m[update.Message.Chat.ID]

			if update.Message.IsCommand() {
				cmdText := update.Message.Command()
				// check
				if cmdText == "menu" {
				} else if cmdText == "start" {
					if a, _ := checkUserChatExist(update.Message.Chat.ID, db); !a {
						flag = 1
						if !ok || i < 4 {
							msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Main menu")
							msg.ReplyMarkup = StartMenuKeyboard
						} else {
							msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Что-то на случай наличия регистрации")
						}
					} else {
						msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Вы зареганы")
						msg.ReplyMarkup = CheckinMenuKeyboard
					}
					bot.Send(msg)
				} else if cmdText == "create" {
					flag = 2
					msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Создани е ивента:")
					bot.Send(msg)
				}
			} else {
				if flag == 1 {
					registration(update, bot, &i, msg, db, &flag)
				} else if flag == 2 {

				} else {
					msg = tgbotapi.NewMessage(update.Message.Chat.ID, "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAA")
					bot.Send(msg)
				}
			}
		}
	}
}
