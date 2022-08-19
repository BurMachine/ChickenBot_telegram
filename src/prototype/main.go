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
	flag1 := 1
	m = make(map[int64]int)
	signMap = make(map[int64]*user)
	createMap = make(map[int64]*events)
	i := 0 // флаг регистрации(4 если все ок)
	//createFlag := 0
	for update := range updates {
		var msg tgbotapi.MessageConfig
		if update.Message != nil {
			//_, ok := m[update.Message.Chat.ID]

			if update.Message.IsCommand() {
				cmdText := update.Message.Command()
				// check
				if cmdText == "menu" {
				} else if cmdText == "start" {
					flag = 0
					flag1 = 1
					if a, _ := checkUserChatExist(update.Message.Chat.ID, db); !a {
						flag = 1
						msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Main menu")
						msg.ReplyMarkup = StartMenuKeyboard
					} else {
						msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Вы зареганы")
						msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
						//bot.Send(msg)
						msg.ReplyMarkup = inKeyboard
					}
					bot.Send(msg)
				} else if cmdText == "create_event_command-for-tgcommand-line" {

				}
			} else {
				if flag == 1 {
					registration(update, bot, &i, msg, db, &flag)
				} else if flag == 2 {
					if update.Message.Text == "Да" {
						creation(update, bot, &flag1, msg, db, &flag)
					} else {
						msg = tgbotapi.NewMessage(update.Message.Chat.ID, "нет ок")
						msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
						bot.Send(msg)
						msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Выберете другую функцию из предложенных")
						msg.ReplyMarkup = inKeyboard
						bot.Send(msg)
						flag = 0
					}
					//if createFlag == 0 {
					//	createFlag++
					//} else if createFlag == 1 {
					//	createFlag++
					//} else if createFlag == 2 {
					//	createFlag++
					//}
				} else {
					msg = tgbotapi.NewMessage(update.Message.Chat.ID, "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAA")
					bot.Send(msg)
				}
			}
		} else if update.CallbackQuery != nil {
			callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
			if _, err := bot.Request(callback); err != nil {
				panic(err)
			}
			if update.CallbackQuery.Data == "create_event" {
				msg = tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Создание ивента\nХотите создать ивент?")
				msg.ReplyMarkup = YesOrNo
				flag1 = 0
				flag = 2
			} else if update.CallbackQuery.Data == "see_all_events" {
				msg = tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Просмотр ивентов")
				flag = 3
			} else if update.CallbackQuery.Data == "Chiken-box" {
				msg = tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Чикен Чикен🐣")
				flag = 4
			}
			bot.Send(msg)
		}
	}
}
