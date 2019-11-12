// Introduction
// Given a number determine whether or not it is valid per the Luhn formula.
//
// The Luhn algorithm is a simple checksum formula used to validate a variety of identification numbers, such as credit card numbers and Canadian Social Insurance Numbers.
//
// The task is to check if a given string is valid.
//
// Validating a Number
// Strings of length 1 or less are not valid. Spaces are allowed in the input, but they should be stripped before checking. All other non-digit characters are disallowed.
//
// Example 1: valid credit card number
// 4539 1488 0343 6467
// The first step of the Luhn algorithm is to double every second digit, starting from the right. We will be doubling
//
// 4_3_ 1_8_ 0_4_ 6_6_
// If doubling the number results in a number greater than 9 then subtract 9 from the product. The results of our doubling:
//
// 8569 2478 0383 3437
// Then sum all of the digits:
//
// 8+5+6+9+2+4+7+8+0+3+8+3+3+4+3+7 = 80
// If the sum is evenly divisible by 10, then the number is valid. This number is valid!
//
// Example 2: invalid credit card number
// 8273 1232 7352 0569
// Double the second digits, starting from the right
//
// 7253 2262 5312 0539
// Sum the digits
//
// 7+2+5+3+2+2+6+2+5+3+1+2+0+5+3+9 = 57
// 57 is not evenly divisible by 10, so this number is not valid.

package main

import (
  "fmt"
)

// There is no prepend in Golang, must reverse string to do
// proper modulus operations

// Concept borrowed from: https://github.com/golang/example/blob/master/stringutil/reverse.go
func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func isLuhnValid(input string) bool {
  var sum int
  var current_index int

  for _, r := range reverse(input) {
    value := int(r - '0')
    // For every other digit, double & subtract 9 if greater than 9
    if current_index % 2 == 1 {
      value = (value * 2)
      if value > 9 {
        value = value - 9
      }
    }
    sum += value

    // Iterate to next current_index
    current_index++
  }

  // Return bool: is sum divisable by 10
  return (sum % 10) == 0

}


func main() {
  fmt.Println(isLuhnValid("4539148803436467"))
  fmt.Println(isLuhnValid("1234"))
}
