package main

import (
	"database/sql"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strconv"
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
		createMap[update.Message.From.ID] = new(events)
		createMap[update.Message.From.ID].state = 0
		log.Println(update.Message.From.UserName, "Пошло создание ивента")
		msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Назовите ваше мероприятие .....")
		msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
		bot.Send(msg)
	} else {
		crMap, ok := createMap[update.Message.From.ID]
		if ok {
			botCreation(crMap, update, bot, msg, flag, db)
		} else {
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Creation mistake or no")
			bot.Send(msg)
		}
	}
}

func botCreation(cr *events, update tgbotapi.Update, bot *tgbotapi.BotAPI, msg tgbotapi.MessageConfig, flag *int, db *sql.DB) {
	if cr.state == 0 {
		cr.name = update.Message.Text
		// без проверок пока
		a, _ := lastEventId(db)
		msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Введите description ивента......")
		cr.uniqueCode = strconv.Itoa(a + 1)
		bot.Send(msg)
		cr.state = 1
	} else if cr.state == 1 {
		cr.description = update.Message.Text
		msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Введите дату начала....")
		bot.Send(msg)
		cr.state = 2
	} else if cr.state == 2 {
		cr.startTime = update.Message.Text
		msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Введите дату окончания....")
		bot.Send(msg)
		cr.state = 3
	} else if cr.state == 3 {
		cr.expiresTime = update.Message.Text
		msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Введите формат онлайн или город....")
		bot.Send(msg)
		cr.state = 4
	} else if cr.state == 4 {
		cr.eType = update.Message.Text
		msg = tgbotapi.NewMessage(update.Message.Chat.ID, "OK")
		bot.Send(msg)
		cr.state = 4
		*flag = 0
		addEvent(cr, db)
		delete(createMap, update.Message.From.ID)
		log.Println(update.Message.From.UserName, "Должно было закончится заполг=нение бд", cr.uniqueCode)
		// выдать сообщение - ссылка для регистрации на ивент
	}

}
