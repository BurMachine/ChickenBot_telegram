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
var user = "postgres"
var password = "test"
var dbname = "postgres"
var sslmode = "disable"

var dbInfo = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", host, port, user, password, dbname, sslmode)

//Collecting data from bot

//Creating users table in database
func createTableUsers() error {

	//Connecting to database
	db, err := sql.Open("postgres", dbInfo)
	if err != nil {
		return err
	}
	defer db.Close()

	//Creating users Table
	if _, err = db.Exec(`CREATE TABLE users(ID SERIAL PRIMARY KEY, LOGIN TEXT, USERNAME TEXT, ROLE INT, CAMPUS TEXT);`); err != nil {
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
	if _, err = db.Exec(`CREATE TABLE events(ID SERIAL PRIMARY KEY, TYPE TEXT, DESCRIPTION TEXT, UNIQUE_CODE TEXT, START_TIME TIMESTAMP, EXPIRIES_TIME TIMESTAMP);`); err != nil {
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
	if _, err = db.Exec(`CREATE TABLE chats(ID SERIAL PRIMARY KEY, CHAT_ID BIGINT);`); err != nil {
		return err
	}

	return nil
}
