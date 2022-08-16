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

	var ucfg tgbotapi.UpdateConfig = tgbotapi.NewUpdate(0) // инициализируем канал, куда будут прилетать обновления от API
	ucfg.Timeout = 60
	updates, err := bot.GetUpdatesChan(ucfg) // считываем обновления из канала

	for update := range updates { // читаем обновления из канала
		if update.Message != nil {
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
			if ParseCommand(update.Message.Text, update, bot) == 0 {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
				msg.ReplyToMessageID = update.Message.MessageID
				bot.Send(msg)
			}
		}
	}
}
