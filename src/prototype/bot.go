package main

import (
	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"log"
)

func (telegramBot *TelegramBot) Init() {
	botAPI, err := tgbotapi.NewBotAPI(TELEGRAM_BOT_API_KEY) // Инициализация API
	if err != nil {
		log.Fatal(err)
	}
	telegramBot.API = botAPI
	botUpdate := tgbotapi.NewUpdate(0) // Инициализация канала обновлений
	botUpdate.Timeout = 64
	botUpdates, err := telegramBot.API.GetUpdatesChan(botUpdate)
	if err != nil {
		log.Fatal(err)
	}
	telegramBot.Updates = botUpdates
}

//func (telegramBot *TelegramBot) Start() {
//	for update := range telegramBot.Updates {
//		if update.Message != nil {
//			// Если сообщение есть  -> начинаем обработку
//			telegramBot.analyzeUpdate(update)
//		}
//	}
//}
//
//// Начало обработки сообщения
//func (telegramBot *TelegramBot) analyzeUpdate(update tgbotapi.Update) {
//	chatID := update.Message.Chat.ID
//	if telegramBot.findUser(chatID) { // Есть ли пользователь в БД?
//		telegramBot.analyzeUser(update)
//	} else {
//		telegramBot.createUser(User{chatID}) // Создаём пользователя
//		telegramBot.requestContact(chatID)   // Запрашиваем номер
//	}
//}
//
