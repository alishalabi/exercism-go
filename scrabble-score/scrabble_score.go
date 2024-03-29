package main

import(
  "fmt"
  "strings"
)

func setMap() map[string]int {
  m := make(map[rune]int)
  // keys should be runes
  // map[int][]rune

  // 1's
  m["A"] = 1
  m["E"] = 1
  m["I"] = 1
  m["O"] = 1
  m["U"] = 1
  m["L"] = 1
  m["N"] = 1
  m["R"] = 1
  m["S"] = 1
  m["T"] = 1

  // 2's
  m["D"] = 2
  m["G"] = 2

  // 3's
  m["B"] = 3
  m["C"] = 3
  m["M"] = 3
  m["P"] = 3

  // 4's
  m["F"] = 4
  m["H"] = 4
  m["V"] = 4
  m["W"] = 4

  // 5's
  m["K"] = 5

  // 8's
  m["J"] = 8
  m["X"] = 8

  // 10's
  m["Q"] = 10
  m["Z"] = 10

  return m
}

func getScore(input string) int {
  score := 0
  m := setMap()
  upperInput := strings.ToUpper(input)
  for _, letter := range upperInput {
    // fmt.Println(string(letter))
    letterScore := m[letter]
    // fmt.Println(letterScore)
    score += letterScore

  }
  return score
}

func main() {
  fmt.Println(getScore("Hello"))
}
