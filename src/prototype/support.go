package main

import (
	"fmt"
	"log"
	"math/rand"
	"regexp"
)

func check_name(name string) bool {
	matched, _ := regexp.MatchString(`^[a-zA-Zа-яА-Я]+$`, name)
	return matched
}

func check_login(name string) bool {
	matched, _ := regexp.MatchString(`^[a-zA-Z]+$`, name)
	return matched
}

// Callback functions

//func callbackQueryProc(update tgbotapi.Update)  {
//
//}

func Hash() string {
	b := make([]byte, 8)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatal(err)
	}
	uuid := fmt.Sprintf("%x%x%x%x%x",
		b[0:2], b[2:3], b[3:4], b[4:5], b[5:])
	fmt.Println("hash", uuid)
	return uuid
}
