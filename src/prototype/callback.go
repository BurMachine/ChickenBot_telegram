package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func CallBackQueryProc(bot *tgbotapi.BotAPI, update tgbotapi.Update, msg tgbotapi.MessageConfig, flag1 *int, flag *int) {
	callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
	if _, err := bot.Request(callback); err != nil {
		panic(err)
	}
	if update.CallbackQuery.Data == "create_event" {
		msg = tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Создание ивента\nХотите создать ивент?")
		msg.ReplyMarkup = YesOrNo
		*flag1 = 0
		*flag = 2
	} else if update.CallbackQuery.Data == "see_all_events" {
		msg = tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Вывести список всех зарегестрированных ивентов?(адм)")
		msg.ReplyMarkup = YesOrNo
		*flag = 3
	} else if update.CallbackQuery.Data == "Chiken-box" {
		msg = tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Вывести список всех чекинов?")
		msg.ReplyMarkup = YesOrNo
		*flag = 4
	} else if update.CallbackQuery.Data == "see_all_events_user" {
		msg = tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "СМОТРЕТЬ ВСЕ ИВЕНТВ ОТ ЛИЦА ЮЗЕРА")
		msg.ReplyMarkup = YesOrNo
		*flag = 3
	} else if update.CallbackQuery.Data == "Chiken-box_user" {
		msg = tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "СHECK USER\nВведите код ивента")
		*flag = 5
	}
	bot.Send(msg)
}
