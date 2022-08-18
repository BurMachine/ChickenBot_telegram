package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
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
func createTableUsers() error {

	//Connecting to database
	db, err := sql.Open("postgres", dbInfo)
	if err != nil {
		return err
	}
	defer db.Close()

	//Creating users Table
	if _, err = db.Exec(`CREATE TABLE IF NOT EXISTS users(ID SERIAL PRIMARY KEY, LOGIN TEXT, USERNAME TEXT, ROLE INT, CAMPUS TEXT);`); err != nil {
		return err
	}

	return nil
}

//Creating events table in database
func createTableEvents() error {

	//Connecting to database
	db, err := sql.Open("postgres", dbInfo)
	if err != nil {
		return err
	}
	defer db.Close()

	//Creating users Table
	if _, err = db.Exec(`CREATE TABLE IF NOT EXISTS events(ID SERIAL PRIMARY KEY, TYPE TEXT, DESCRIPTION TEXT, UNIQUE_CODE TEXT, START_TIME TIMESTAMP, EXPIRIES_TIME TIMESTAMP);`); err != nil {
		return err
	}

	return nil
}

//Creating chats table in database
func createTableChats() error {

	//Connecting to database
	db, err := sql.Open("postgres", dbInfo)
	if err != nil {
		return err
	}
	defer db.Close()

	//Creating chats Table
	if _, err = db.Exec(`CREATE TABLE IF NOT EXISTS chats(ID SERIAL PRIMARY KEY, CHAT_ID BIGINT);`); err != nil {
		return err
	}

	return nil
}

//Add record in users table
func addUser(us *user) error {
	//Connecting to database
	db, err := sql.Open("postgres", dbInfo)
	if err != nil {
		return err
	}
	defer db.Close()

	if _, err = db.Exec("INSERT INTO users (login, username, role, campus) values($1, $2, $3, $4)", us.login, us.name, 0, us.campus); err != nil {
		return err
	}

	return nil
}

// Add new event in DB
func addEvent(event *events) error {
	//Connecting to database
	db, err := sql.Open("postgres", dbInfo)
	if err != nil {
		return err
	}
	defer db.Close()

	if _, err = db.Exec("INSERT INTO event(TYPE, DESCRIPTION, UNIQUE_CODE, START_TIME, EXPIRIES_TIME) values('$1','$2','$3','$4','$5')"); err != nil {
		return err
	}

	return nil
}

//Check user in DB
func checkUserExist(name string) (int, error) {

	var count int
	//Connecting to database
	db, err := sql.Open("postgres", dbInfo)
	if err != nil {
		return 0, err
	}
	defer db.Close()

	//Counting number of users
	row := db.QueryRow("SELECT * FROM users WHERE username = '$1'")
	err = row.Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}
