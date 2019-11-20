package main

import (
	"fmt"
)

func isLeapYear(year int) bool {
	if year % 400 == 0 {
		return true
	} else if year % 100 == 0 {
		return false
	} else if year % 4 == 0 {
		return true
	}
	return false
}

func main() {
	fmt.Println(isLeapYear(1997))
	fmt.Println(isLeapYear(2000))
	fmt.Println(isLeapYear(1996))
}
