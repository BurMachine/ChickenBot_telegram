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

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)
	flag := 0
	m = make(map[int64]int)
	signMap = make(map[int64]*user)
	i := 0 // флаг регистрации(4 если все ок)
	for update := range updates {
		var msg tgbotapi.MessageConfig
		if update.Message != nil {
			_, ok := m[update.Message.Chat.ID]

			if update.Message.IsCommand() {
				cmdText := update.Message.Command()
				if cmdText == "start" {
				} else if cmdText == "menu" {
					flag = 1
					if !ok || i < 4 {
						msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Main menu")
						msg.ReplyMarkup = StartMenuKeyboard
					} else {
						msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Что-то на случай наличия регистрации")
					}
					bot.Send(msg)
				} else if cmdText == "gorod" {
					flag = 2
					msg = tgbotapi.NewMessage(update.Message.Chat.ID, "fbawiuyfgaoisugfawiug")
					msg.ReplyMarkup = CampusMenuKeyboard
					bot.Send(msg)
				}
			} else {
				if flag == 1 {
					if update.Message.Text == StartMenuKeyboard.Keyboard[0][0].Text && i == 0 {
						i++
						signMap[update.Message.From.ID] = new(user)
						signMap[update.Message.From.ID].state = 0
						log.Println(update.Message.From.UserName, update.Message.Text)
						msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Как вас зовут?")
						bot.Send(msg)
						log.Print("i=", i)
					} else {
						us, ok := signMap[update.Message.From.ID]
						log.Print(flag)
						if ok {
							botReg(us, update, bot, msg, &i, db)
							log.Print(us, flag)
						} else {
							msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Я вас не понял..(мб зареган)")
							bot.Send(msg)
						}
						if i == 4 {
							//i = 0
							flag = 0
						}
					}
				} else if flag == 2 {
					if update.Message.Text == CampusMenuKeyboard.Keyboard[0][0].Text {
						msg = tgbotapi.NewMessage(update.Message.Chat.ID, "1")
						bot.Send(msg)
					} else if update.Message.Text == CampusMenuKeyboard.Keyboard[0][1].Text {
						msg = tgbotapi.NewMessage(update.Message.Chat.ID, "2")
						bot.Send(msg)
					} else if update.Message.Text == CampusMenuKeyboard.Keyboard[0][2].Text {
						msg = tgbotapi.NewMessage(update.Message.Chat.ID, "3")
						bot.Send(msg)
					} else {
						msg = tgbotapi.NewMessage(update.Message.Chat.ID, "123")
						bot.Send(msg)
					}
				} else {
					msg = tgbotapi.NewMessage(update.Message.Chat.ID, "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAA")
					bot.Send(msg)
				}
			}
		}
	}
}
