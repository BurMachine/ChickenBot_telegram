package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

//var host = os.Getenv("HOST")
//var port = os.Getenv("PORT")
//var user = os.Getenv("USER")
//var password = os.Getenv("PASSWORD")
//var dbname = os.Getenv("DBNAME")
//var sslmode = os.Getenv("SSLMODE")

var host = "localhost"
var port = "5432"
var user_db = "postgres"
var password = "test"
var dbname = "postgres"
var sslmode = "disable"

var dbInfo = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", host, port, user_db, password, dbname, sslmode)

//Creating users table in database
func openDatabase() *sql.DB {

	//Connecting to database
	db, err := sql.Open("postgres", dbInfo)
	if err != nil {
		log.Panic(err)
	}

	//Creating DB Tables
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users(ID SERIAL PRIMARY KEY, chatID BIGINT, LOGIN TEXT, USERNAME TEXT, ROLE INT, CAMPUS TEXT);`)
	if err != nil {
		log.Panic(err)
	}
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS events(ID SERIAL PRIMARY KEY, eType TEXT, name TEXT, description TEXT, uniqueCode TEXT, startTime TEXT, expiriesTime TEXT);`)
	if err != nil {
		log.Panic(err)
	}
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS checkins(ID SERIAL PRIMARY KEY, login TEXT, uniqueCode TEXT);`)
	if err != nil {
		log.Panic(err)
	}
	return db
}

//Add record in users table
func addUser(us *user, db *sql.DB) error {
	if _, err := db.Exec("INSERT INTO users (chatID, login, username, role, campus) values($1, $2, $3, $4, $5);",
		us.chatID, us.login, us.name, 0, us.campus); err != nil {
		return err
	}
	return nil
}

// Add new event in DB
func addEvent(event *events, db *sql.DB) error {
	if _, err := db.Exec("INSERT INTO events (etype, name, description, uniqueCode, startTime, expiriesTime) values($1, $2, $3, $4, $5, $6);",
		event.eType,
		event.name,
		event.description,
		event.uniqueCode,
		event.startTime,
		event.expiresTime); err != nil {
		return err
	}
	return nil
}

//Check user in DB
func checkUserNameExist(login string, db *sql.DB) (bool, error) {
	//Counting number of users
	var count int
	row := db.QueryRow("SELECT COUNT(DISTINCT username) FROM users WHERE username = $1;", login)
	err := row.Scan(&count)
	if err != nil {
		return false, err
	}
	if count > 0 {
		return true, nil
	}
	return false, nil
}

/*
	return true if user exist
*/
func checkUserChatExist(chatID int64, db *sql.DB) (bool, error) {
	var count int64
	row := db.QueryRow("SELECT COUNT(DISTINCT username) FROM users WHERE chatid = $1;", chatID)
	err := row.Scan(&count)
	if err != nil {
		return false, err
	}
	if count > 0 {
		return true, nil
	}
	return false, nil
}

func lastEventId(db *sql.DB) (int, error) {
	var count int
	row := db.QueryRow("SELECT MAX(id) FROM events;")
	err := row.Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func checkUserCheckin(chatID int64, db *sql.DB) (bool, error) {
	var login string
	var count int64
	row := db.QueryRow("SELECT login FROM users WHERE chatid = $1;", chatID)
	err := row.Scan(&login)
	if err != nil {
		return false, err
	}
	row = db.QueryRow("SELECT COUNT(DISTINCT login) FROM checkins WHERE login = $1;", login)
	err = row.Scan(&count)
	if err != nil {
		return false, err
	}
	if count > 0 {
		return true, nil
	}
	return false, nil
}

// Add user checkin in DB
func addCheckin(chatID int64, event *events, db *sql.DB) error {
	var login string
	row := db.QueryRow("SELECT login FROM users WHERE chatid = $1;", chatID)
	err := row.Scan(&login)
	if err != nil {
		return err
	}
	if _, err := db.Exec("INSERT INTO checkins (login, uniqueCode) values($1, $2);",
		login,
		event.uniqueCode); err != nil {
		return err
	}
	return nil
}

func checkEventExist(uniqueCode string, db *sql.DB) (bool, error) {
	var count int64
	row := db.QueryRow("SELECT COUNT(DISTINCT uniqueCode) FROM events WHERE uniqueCode = $1;", uniqueCode)
	err := row.Scan(&count)
	if err != nil {
		return false, err
	}
	if count > 0 {
		return true, nil
	}
	return false, nil
}

func isUserAdmin(chatID int64, db *sql.DB) (bool, error) {
	var role int
	row := db.QueryRow("SELECT role FROM users WHERE chatid = $1;", chatID)
	err := row.Scan(&role)
	if err != nil {
		return false, err
	}
	if role == 1 {
		return true, nil
	}
	return false, nil
}

func outputAllEvents(db *sql.DB) error {
	var ctx context.Context
	rows, err := db.QueryContext(ctx, "SELECT eType, name, description, startTime, expiriesTime FROM events;")
	if err != nil {
		return err
	}
	for rows.Next() {
		var eType, name, description, startTime, expiriesTime string
		if err := rows.Scan(&eType, &name, &description, &startTime, &expiriesTime); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Название события: %s \nТип: %s\nОписание %s\n Начало: %s\nОкончание: %s", name, eType, description, startTime, expiriesTime)
	}
	return nil
}
