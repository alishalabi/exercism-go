// Introduction
// Bob is a lackadaisical teenager. In conversation, his responses are very limited.
//
// Bob answers 'Sure.' if you ask him a question, such as "How are you?".
//
// He answers 'Whoa, chill out!' if you YELL AT HIM (in all capitals).
//
// He answers 'Calm down, I know what I'm doing!' if you yell a question at him.
//
// He says 'Fine. Be that way!' if you address him without actually saying anything.
//
// He answers 'Whatever.' to anything else.
//
// Bob's conversational partner is a purist when it comes to written communication and always follows normal rules regarding sentence punctuation in English.


package main

import (
	"unicode"
	"fmt"
)

func stringIsUpper(input string) bool {
	r := []rune(input)

	for i := 0; i < len(r); i++ {
		ch := r[i]
		if unicode.IsLetter(ch) {
			if unicode.IsUpper(ch) == false {
				return false
			}
		}

	}
	return true
}

func Speak(input string) string {
	if input == "How are you?" {
		return "Sure."
	} else if input == ""{
		return "Fine. Be that way!"
	} else if stringIsUpper(input) && input[len(input) - 1:] == "?"{
		return "Calm down, I know what I'm doing!"
	} else if stringIsUpper(input) {
		return "Whoa, chill out!"
	} else {
		return "Whatever."
	}
}

func main() {
	fmt.Println(Speak("How are you?"))
	fmt.Println(Speak(""))
	fmt.Println(Speak("HELLO?"))
	fmt.Println(Speak("HELLO"))
	fmt.Println(Speak("What's poppin?"))
}
