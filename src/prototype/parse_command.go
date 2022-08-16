package main

import (
	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"log"
)

func ParseCommand(str string, id tgbotapi.Update, bot *tgbotapi.BotAPI) (a bool) {
	if str == "/help" {
		msg := tgbotapi.NewMessage(id.Message.Chat.ID, "Здарова")
		msg1, err := bot.Send(msg)
		if err != nil {
			log.Println(err, msg1)
			return false
		}
	} else {
		return false
	}
	return true
}
