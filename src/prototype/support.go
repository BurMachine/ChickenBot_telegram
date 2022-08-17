package main

import "regexp"

func check_name(name string) bool {
	matched, _ := regexp.MatchString(`^[a-zA-Zа-яА-Я]+$`, name)
	return matched
}

func check_login(name string) bool {
	matched, _ := regexp.MatchString(`^[a-zA-Z]+$`, name)
	return matched
}
