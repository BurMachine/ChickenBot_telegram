package main

import tgbotapi "github.com/Syfaro/telegram-bot-api"

func botReg(us *user, update tgbotapi.Update, bot *tgbotapi.BotAPI, msg tgbotapi.MessageConfig) {
	if us.state == 0 {
		us.name = update.Message.Text
		msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Какой у вас логин на платформе?")
		bot.Send(msg)
		us.state = 1
	} else if us.state == 1 {
		us.login = update.Message.Text
		msg = tgbotapi.NewMessage(update.Message.Chat.ID, "С какого вы кампуса?")
		msg.ReplyMarkup = CampusMenuKeyboard
		bot.Send(msg)
		us.state = 2
	} else if us.state == 2 {
		us.campus = update.Message.Text
		msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Окей, запомнил!")
		msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
		bot.Send(msg)
		delete(signMap, update.Message.From.ID)
	}
}
