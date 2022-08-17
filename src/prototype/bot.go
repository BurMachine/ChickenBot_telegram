package main

import tgbotapi "github.com/Syfaro/telegram-bot-api"

func botReg(us *user, update tgbotapi.Update, bot *tgbotapi.BotAPI, msg tgbotapi.MessageConfig) {
	if us.state == 0 {
		us.name = update.Message.Text
		if !check_name(us.name) {
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Некорректное имя - используйте только буквы")
			bot.Send(msg)
		} else {
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Какой у вас логин на платформе?")
			bot.Send(msg)
			us.state = 1
		}
	} else if us.state == 1 {
		us.login = update.Message.Text
		if !check_login(us.login) {
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Некорректный логин, используйте только латиницу")
			bot.Send(msg)
		} else {
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, "С какого вы кампуса?")
			msg.ReplyMarkup = CampusMenuKeyboard
			bot.Send(msg)
			us.state = 2
		}
	} else if us.state == 2 {
		us.campus = update.Message.Text
		if us.campus != "Казань" && us.campus != "Москва" && us.campus != "Новосибирск" {
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Некорректный город")
			bot.Send(msg)
		} else {
			// addUser(us)
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Окей, запомнил!")
			msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
			bot.Send(msg)
			delete(signMap, update.Message.From.ID)

		}
	}
}
