package main

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

const (
	TELEGRAM_BOT_API_KEY       = "5775513785:AAGy6Ht6IYgaZUVfLOmyyYiviwtJfJhmKu8" // API  ключ, который мы получили у BotFather
	POSTGRESS_CONNECTION_URL   = "localhost"                                      // Адрес сервера PostgressDB
	POSTGRESS_DATABASE_NAME    = "regbot"                                         // Название базы данных
	POSTGRESS_COLLECTION_USERS = "users"                                          // Название таблицы
)

var ArrEmojiText []string = []string{"💣", "📸", "📟", "✈", "🚀",
	"🛸", "🍾", "☕", "🍕", "🥑",
	"🦖", "🦉", "🐣", "🦩", "🦁",
	"🐈", "🦄", "🐅", "🦣", "☠",
	"🤬", "😈", "🌠", "🪐", "🔥",
	"🌈", "🌝", "💎", "🗿", "🦊",
	"👾", "👻", "💩", "🤡", "🤖",
	"👽", "🔑", "💰", "📱", "🕶",
	"🥽", "👑", "🎓", "🎨", "🎮",
	"🪄", "⚡", "🦝", "☁️", "⭐️"}

type User struct {
	Chat_ID int64
}

type TelegramBot struct {
	API                   *tgbotapi.BotAPI        // API телеграмма
	Updates               tgbotapi.UpdatesChannel // Канал обновлений
	ActiveContactRequests []int64                 // ID чатов, от которых мы ожидаем номер
}

type user struct {
	chatID int64
	state  int // 0, 1, 2, 3
	name   string
	login  string
	campus string
	role   int
}

type events struct {
	state       int
	name        string
	eType       string
	description string
	uniqueCode  string
	startTime   string //проверить типб в БД timestamp
	expiresTime string //проверить типб в БД timestamp
}

var signMap map[int64]*user
var createMap map[int]*events

var StartMenuKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Регистрация"),
	),
)

var CampusMenuKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Казань"),
		tgbotapi.NewKeyboardButton("Москва"),
		tgbotapi.NewKeyboardButton("Новосибирск"),
	),
)

var CheckinMenuKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Создать"),
		tgbotapi.NewKeyboardButton("Чекин"),
		tgbotapi.NewKeyboardButton("123"),
	),
)
var inKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData(ArrEmojiText[4]+"Создать ивент"+ArrEmojiText[4], "create_event"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData(ArrEmojiText[8]+"Просмотр всех ивентов"+ArrEmojiText[8], "see_all_events"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData(ArrEmojiText[12]+"Чикен"+ArrEmojiText[12], "Chiken-box"),
	),
)
