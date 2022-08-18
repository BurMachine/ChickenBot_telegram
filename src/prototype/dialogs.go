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
		log.Print("i=", i)
	} else {
		us, ok := signMap[update.Message.From.ID]
		//log.Print(flag)
		if ok {
			botReg(us, update, bot, msg, i, db)
			//log.Print(us, flag)
		} else {
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Я вас не понял..(мб зареган)")
			bot.Send(msg)
		}
		if *i == 4 {
			//i = 0
			*flag = 0
		}

	}
}
