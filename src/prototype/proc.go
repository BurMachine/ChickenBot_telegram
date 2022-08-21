package main

import (
	"database/sql"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strings"
)

func commandProc(update tgbotapi.Update, bot *tgbotapi.BotAPI, db *sql.DB, msg tgbotapi.MessageConfig, flag1 *int, flag *int) {
	cmdText := update.Message.Command()
	if checkDeepLink(update.Message.Text) {
		code := strings.Split(update.Message.Text, " ")
		log.Println(code[1])
		checkin(update, bot, msg, flag, db, code[1])
	} else if cmdText == "start" {
		log.Println(update.Message.Text, update.Message.Chat.UserName)
		*flag = 0
		*flag1 = 1
		if a, _ := checkUserChatExist(update.Message.Chat.ID, db); !a {
			*flag = 1
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Main menu")
			msg.ReplyMarkup = StartMenuKeyboard
		} else {
			if a, _ := isUserAdmin(update.Message.Chat.ID, db); a {
				msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Вы зареганы как админ")
				msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
				msg.ReplyMarkup = inKeyboard
			} else {
				msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Вы зареганы как юзер")
				msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
				msg.ReplyMarkup = inKeyboard_user
			}
		}
		bot.Send(msg)
	}

}

func nonCommandProc(update tgbotapi.Update, bot *tgbotapi.BotAPI, db *sql.DB, msg tgbotapi.MessageConfig, flag1 *int, flag *int, i *int) {
	if *flag == 1 {
		registration(update, bot, i, msg, db, flag)
	} else if *flag == 2 {
		if update.Message.Text == "Да" || *flag1 == 3 {
			creation(update, bot, flag1, msg, db, flag)
		} else {
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Принято!")
			msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
			bot.Send(msg)
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Выберете другую функцию из предложенных")
			msg.ReplyMarkup = inKeyboard
			bot.Send(msg)
			*flag = 0
		}
	} else if *flag == 3 {
		msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Принято!!")
		msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
		bot.Send(msg)
		if update.Message.Text == "Да" {
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, "печать списка ивентов из базы")
			outputAllEvents(db, update, bot)
			bot.Send(msg)
		} else {
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Выберете другую функцию из предложенных")
			if a, _ := isUserAdmin(update.Message.Chat.ID, db); a {
				msg.ReplyMarkup = inKeyboard
			} else {
				msg.ReplyMarkup = inKeyboard_user
			}
			bot.Send(msg)
			*flag = 0
		}
	} else if *flag == 5 {
		code := update.Message.Text
		checkin(update, bot, msg, flag, db, code)
	} else if *flag == 4 {
		if update.Message.Text == "Да" {
			outputAllCheckins(db, update, bot)
		} else {
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Не вышло выкатить список чекинов!!!!!!!")
			msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
			bot.Send(msg)
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Выберете другую функцию из предложенных")
			msg.ReplyMarkup = inKeyboard
			*flag = 0
		}
		bot.Send(msg)
	} else {
		msg = tgbotapi.NewMessage(update.Message.Chat.ID, "А вот сейчас не понял.")
		msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
		bot.Send(msg)
	}
}
