package main

import (
	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"log"
)

func main() {
	bot, err := tgbotapi.NewBotAPI(TELEGRAM_BOT_API_KEY) // подключаемся к боту с помощью токена
	if err != nil {
		log.Panic(err)
	}

	//bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	// Loop through each update.
	for update := range updates {
		var msg tgbotapi.MessageConfig
		if update.Message != nil {
			if update.Message.IsCommand() {
				cmdText := update.Message.Command()
				if cmdText == "start" {
				} else if cmdText == "menu" {
					msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Main menu")
					msg.ReplyMarkup = StartMenuKeyboard
					bot.Send(msg)
				}
			} else {
				if update.Message.Text == StartMenuKeyboard.Keyboard[0][0].Text {
					signMap[update.Message.From.ID] = new(user)
					signMap[update.Message.From.ID].state = 0
					log.Println(update.Message.From.UserName, update.Message.Text)
					msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Как вас зовут?")
					bot.Send(msg)
				} else {
					us, ok := signMap[update.Message.From.ID]
					if ok {
						botReg(us, update, bot, msg)
						log.Print(us)
					} else {
						msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Я вас не понял..")
						bot.Send(msg)
					}
				}
			}
			//switch update.Message.Text {
			//case "/start":
			//	msg = start(update)
			//	//bot.Send(msg)
			//case "/help":
			//	msg = CloseStartMenu(update, msg)
			//	msg = help(update)
			//case "Регистрация":
			//	msg = RegUser(update)
			//}
			//if RegFlag == 1 || RegFlag == 0 {
			//	msg = StartMenu(update, &RegFlag)
			//}
			//_, err = bot.Send(msg)
		}
	}
}

func init() {
	signMap = make(map[int]*user)
}
