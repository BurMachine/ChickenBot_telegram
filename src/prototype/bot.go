package main

import (
	"database/sql"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/skip2/go-qrcode"
	"io/ioutil"
	"log"
	"strconv"
	time2 "time"
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

func botCreation(cr *events, update tgbotapi.Update, bot *tgbotapi.BotAPI, msg tgbotapi.MessageConfig, flag *int, db *sql.DB, flag1 *int) {
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
		time_str := strconv.Itoa(time2.Now().Nanosecond())
		a := Hash() + time_str
		c := "https://t.me/evcheckinbot?start=" + a

		cr.eType = update.Message.Text
		msg = tgbotapi.NewMessage(update.Message.Chat.ID, "OK\n"+c)
		err := qrcode.WriteFile(a, qrcode.Medium, 256, "qr.png")
		if err != nil {
			log.Println(err, "перевод в qr")
		}
		data, _ := ioutil.ReadFile("qr.png")
		b := tgbotapi.FileBytes{Name: "qr.png", Bytes: data}
		msg := tgbotapi.NewPhoto(update.Message.Chat.ID, b)
		msg.Caption = "QR-код для чекина"
		if z, _ := isUserAdmin(update.Message.Chat.ID, db); z {
			msg.ReplyMarkup = inKeyboard
		} else {
			msg.ReplyMarkup = inKeyboard_user
		}
		bot.Send(msg)
		cr.state = 4
		*flag = 0
		*flag1 = 1
		cr.uniqueCode = a
		addEvent(cr, db)

		delete(createMap, update.Message.From.ID)
		log.Println(update.Message.From.UserName, "Должно было закончится заполнение бд", cr.uniqueCode)
		// выдать сообщение - ссылка для регистрации на ивент
	}

}
