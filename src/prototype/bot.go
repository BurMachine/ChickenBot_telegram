package main

import (
	"database/sql"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func botReg(us *user, update tgbotapi.Update, bot *tgbotapi.BotAPI, msg tgbotapi.MessageConfig, i *int, db *sql.DB, flag *int) {
	if us.state == 0 && *i == 1 {
		us.name = update.Message.Text
		if !check_name(us.name) {
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Некорректное имя - используйте только буквы")
			bot.Send(msg)
		} else {
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Какой у вас логин на платформе?")
			us.chatID = update.Message.Chat.ID
			bot.Send(msg)
			us.state = 1
			*i++
		}
	} else if us.state == 1 && *i == 2 {
		us.login = update.Message.Text
		if !check_login(us.login) {
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Некорректный логин, используйте только латиницу")
			bot.Send(msg)
		} else {
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, "С какого вы кампуса?")
			msg.ReplyMarkup = CampusMenuKeyboard
			bot.Send(msg)
			us.state = 2
			*i++
		}
	} else if us.state == 2 && *i == 3 {
		us.campus = update.Message.Text
		if us.campus != "Казань" && us.campus != "Москва" && us.campus != "Новосибирск" {
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Некорректный город")
			bot.Send(msg)
		} else {
			// addUser(us)

			msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Окей, запомнил!")
			msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
			bot.Send(msg)
			// check user exists
			addUser(us, db)
			delete(signMap, update.Message.From.ID)
			*i = 0
			*flag = 0
		}
	}
}

func botCreation(cr *events, update tgbotapi.Update, bot *tgbotapi.BotAPI, msg tgbotapi.MessageConfig, flag1 *int, db *sql.DB) {
	if *flag1 == 1 {
		cr.name = 
	}

}
