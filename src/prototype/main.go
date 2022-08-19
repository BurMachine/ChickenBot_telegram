package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

var m map[int64]int

func main() {
	bot, err := tgbotapi.NewBotAPI("5775513785:AAGy6Ht6IYgaZUVfLOmyyYiviwtJfJhmKu8") // подключаемся к боту с помощью токена
	if err != nil {
		log.Panic(err)
	}
	db := openDatabase()
	//bot.Debug = true
	//db, err := sql.Open("postgres", dbInfo)

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)
	flag := 0
	flag1 := 1
	//check_in := 0
	m = make(map[int64]int)
	signMap = make(map[int64]*user)
	createMap = make(map[int64]*events)
	i := 0 // флаг регистрации(4 если все ок)
	//createFlag := 0
	for update := range updates {
		var msg tgbotapi.MessageConfig
		if update.Message != nil {
			//_, ok := m[update.Message.Chat.ID]

			if update.Message.IsCommand() {
				cmdText := update.Message.Command()
				// check
				if cmdText == "start " {

				} else if cmdText == "start" {
					log.Println(update.Message.Text, update.Message.Chat.UserName)
					flag = 0
					flag1 = 1
					if a, _ := checkUserChatExist(update.Message.Chat.ID, db); !a {
						flag = 1
						msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Main menu")
						msg.ReplyMarkup = StartMenuKeyboard
					} else {
						if a, _ := isUserAdmin(update.Message.Chat.ID, db); a {
							msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Вы зареганы как админ")
							msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
							//bot.Send(msg)
							msg.ReplyMarkup = inKeyboard
						} else {
							msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Вы зареганы как юзер")
							msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
							//bot.Send(msg)
							msg.ReplyMarkup = inKeyboard_user
						}
					}
					bot.Send(msg)
				} else if cmdText == "create_event_command-for-tgcommand-line" {

				}
			} else {
				if flag == 1 {
					registration(update, bot, &i, msg, db, &flag)
				} else if flag == 2 {
					if update.Message.Text == "Да" || flag1 == 3 {
						creation(update, bot, &flag1, msg, db, &flag)
					} else {
						msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Принято!")
						msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
						bot.Send(msg)
						msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Выберете другую функцию из предложенных")
						msg.ReplyMarkup = inKeyboard
						bot.Send(msg)
						flag = 0
					}
					//if createFlag == 0 {
					//	createFlag++
					//} else if createFlag == 1 {
					//	createFlag++
					//} else if createFlag == 2 {
					//	createFlag++
					//}
				} else if flag == 3 {
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
						flag = 0
					}
				} else if flag == 5 {
					// Чекин
					if a, err := checkUserCheckin(update.Message.Chat.ID, update.Message.Text, db); a && err == nil {
						// пишем что он уж внесен в базу
						msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Уже заСHECKINен")
						bot.Send(msg)
					} else {
						if a, _ := checkUserChatExist(update.Message.Chat.ID, db); a {
							// добавлем его в базу
							if a, err := checkEventExist(update.Message.Text, db); a && err == nil { // отсюда
								err = addCheckin(update.Message.Chat.ID, update.Message.Text, db)
								if err != nil {
									log.Print(err, 123123123)
									msg = tgbotapi.NewMessage(update.Message.Chat.ID, "СHECKIN failed")
								} else {
									msg = tgbotapi.NewMessage(update.Message.Chat.ID, "СHECKIN прошел успешно!!!!!!!!!!!!!!!!!!!!!!!")
								}
							} else { // до сюда
								msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Ивента с таким кодом не существует")
							}
							bot.Send(msg)
						} else {
							msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Сначала зарегистрируйтесь, потом вернитесь по ссылке")
							msg.ReplyMarkup = StartMenuKeyboard
							bot.Send(msg)
							flag = 1
						}
					}
				} else {
					msg = tgbotapi.NewMessage(update.Message.Chat.ID, "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAA")
					bot.Send(msg)
				}
			}
		} else if update.CallbackQuery != nil {
			callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
			if _, err := bot.Request(callback); err != nil {
				panic(err)
			}
			if update.CallbackQuery.Data == "create_event" {
				msg = tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Создание ивента\nХотите создать ивент?")
				msg.ReplyMarkup = YesOrNo
				flag1 = 0
				flag = 2
			} else if update.CallbackQuery.Data == "see_all_events" {
				msg = tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Вывести список всех зарегестрированных ивентов?(адм)")
				msg.ReplyMarkup = YesOrNo
				flag = 3
			} else if update.CallbackQuery.Data == "Chiken-box" {
				msg = tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Чикен Чикен🐣")
				flag = 4
			} else if update.CallbackQuery.Data == "see_all_events_user" {
				msg = tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "СМОТРЕТЬ ВСЕ ИВЕНТВ ОТ ЛИЦА ЮЗЕРА")
				msg.ReplyMarkup = YesOrNo
				flag = 3
			} else if update.CallbackQuery.Data == "Chiken-box_user" {
				msg = tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "СHECK USER\nВведите код ивента")
				//msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Введите код ивента")
				flag = 5
			}
			bot.Send(msg)
		}
	}
}
