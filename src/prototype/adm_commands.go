package main

import tgbotapi "github.com/Syfaro/telegram-bot-api"

func createEvent(id tgbotapi.Update, bot *tgbotapi.BotAPI) error {
	requestContactMessage := tgbotapi.NewMessage(id.Message.Chat.ID, "В каком городе планируется ивент?")
	KazanButton := tgbotapi.NewKeyboardButton("Казань")
	MoscowButton := tgbotapi.NewKeyboardButton("Москва")
	NovosibirskButton := tgbotapi.NewKeyboardButton("Новосибирск")

	requestContactReplyKeyboard := tgbotapi.NewReplyKeyboard([]tgbotapi.KeyboardButton{KazanButton, MoscowButton, NovosibirskButton})
	requestContactMessage.ReplyMarkup = requestContactReplyKeyboard
	_, err := bot.Send(requestContactMessage)
	if err != nil {
		return err
	}
	return err
}

func regUser(id tgbotapi.Update, bot *tgbotapi.BotAPI) {
	requestContactMessage := tgbotapi.NewMessage(id.Message.Chat.ID, "В каком кампусе ты учишься?")
	KazanButton := tgbotapi.NewKeyboardButton("Казань")
	MoscowButton := tgbotapi.NewKeyboardButton("Москва")
	NovosibirskButton := tgbotapi.NewKeyboardButton("Новосибирск")
	requestContactReplyKeyboard := tgbotapi.NewReplyKeyboard([]tgbotapi.KeyboardButton{KazanButton, MoscowButton, NovosibirskButton})
	requestContactMessage.ReplyMarkup = requestContactReplyKeyboard
	//str := requestContactMessage.Text

	return
}
