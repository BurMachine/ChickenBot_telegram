package main

import (
	"database/sql"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func registration(update tgbotapi.Update, bot *tgbotapi.BotAPI, i *int, msg tgbotapi.MessageConfig, db *sql.DB, flag *int) {
	if update.Message.Text == StartMenuKeyboard.Keyboard[0][0].Text && *i == 0 {
		*i++
		signMap[update.Message.From.ID] = new(user)
		signMap[update.Message.From.ID].state = 0
		log.Println(update.Message.From.UserName, update.Message.Text)
		msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Как вас зовут?")
		bot.Send(msg)
	} else {
		us, ok := signMap[update.Message.From.ID]
		if ok {
			botReg(us, update, bot, msg, i, db, flag)
			//log.Print(us, flag)
		} else {
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Я вас не понял..(мб зареган)")
			bot.Send(msg)
		}
	}
}

func creation(update tgbotapi.Update, bot *tgbotapi.BotAPI, flag1 *int, msg tgbotapi.MessageConfig, db *sql.DB, flag *int) {
	if *flag1 == 0 {
		*flag1 = 1
		// получение id и работа с ним до полного заполнения структуры
		createMap[0] = new(events)
		createMap[0].state = 0
		log.Println(update.Message.From.UserName, "Пошло создание ивента")
		msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Назовите ваше мероприятие .....")
		bot.Send(msg)
	} else {
		crMap, ok := createMap[0]
		if ok {
			botCreation(crMap, update, bot, msg, flag1, db)
		} else {
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Creation mistake or no")
			bot.Send(msg)
		}
	}
}

func botCreation(cr *events, update tgbotapi.Update, bot *tgbotapi.BotAPI, msg tgbotapi.MessageConfig, flag1 *int, db *sql.DB) {
	if cr.state == 0 {
		cr.name = update.Message.Text
		// без проверок пока

	}

}
