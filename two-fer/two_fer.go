package main

import (
	"fmt"
)

func twoFer(name string) {
	if name == "" {
		fmt.Println("One for you, one for me.")
	} else {
		fmt.Printf("One for %v, one for me.", name)
	}
}

func main() {
	twoFer("")
	twoFer("Ali")
}
