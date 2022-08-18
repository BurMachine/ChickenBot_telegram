package main

import (
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

	//Creating users Table
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users(ID SERIAL PRIMARY KEY, chatID INT64, LOGIN TEXT, USERNAME TEXT, ROLE INT, CAMPUS TEXT);`)
	if err != nil {
		log.Panic(err)
	}
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS events(ID SERIAL PRIMARY KEY, TYPE TEXT, DESCRIPTION TEXT, "+
							"UNIQUE_CODE TEXT, START_TIME TIMESTAMP, EXPIRIES_TIME TIMESTAMP);`)
	if err != nil {
		log.Panic(err)
	}
	return db
}

//Add record in users table
func addUser(us *user, db *sql.DB) error {
	if _, err := db.Exec("INSERT INTO users (chatID, login, username, role, campus) values($1, $2, $3, $4, $5)",
		us.chatID, us.login, us.name, 0, us.campus); err != nil {
		return err
	}
	return nil
}

// Add new event in DB
func addEvent(event *events, db *sql.DB) error {
	if _, err := db.Exec("INSERT INTO event(TYPE, DESCRIPTION, UNIQUE_CODE, START_TIME, EXPIRIES_TIME) values('$1','$2','$3','$4','$5')",
		event.eType,
		event.description,
		event.uniqueCode,
		event.startTime,
		event.expiresTime); err != nil {
		return err
	}
	return nil
}

//Check user in DB
func checkUserExist(login string, db *sql.DB) (int, error) {
	//Counting number of users
	var count = 0
	row := db.QueryRow("SELECT * FROM users WHERE username = '$1'", login)
	err := row.Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}
