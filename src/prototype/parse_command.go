package main

import (
	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"log"
)

func ParseCommand(str string, id tgbotapi.Update, bot *tgbotapi.BotAPI) (a int) {
	if str == "/help" {
		msg := tgbotapi.NewMessage(id.Message.Chat.ID, `Краткая справка:
		/registration  - регистрация через школьную почту
		Если в вы зарегистрированы:
		/reg - регистрация на мероприятие из возможных
		/info  - информация о мероприятие
		/event - список мероприятий
		Возможности админа:
		/create - добавить мероприятие
		/delete - удалить мероприятие
		/add_adm - добавить админа
		/delete_adm - удалить админа
		`)
		msg1, err := bot.Send(msg)
		if err != nil {
			log.Println(err, msg1)
			return 0
		}
	} else if str == "/create" {
		err := createEvent(id, bot)
		if err != nil {
			log.Println(err)
			return 0
		}
	} else if str == "/event" {
		msg := tgbotapi.NewMessage(id.Message.Chat.ID, `besligaseliughaosehgoseihg;asodhgo'qn'`)
		msg1, err := bot.Send(msg)
		if err != nil {
			log.Println(err, msg1)
			return 0
		}
	} else {
		return 0
	}
	return 2
}
