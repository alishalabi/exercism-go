package main

type OldSystem struct {

}

func generateNewMap() map[string]int {
  m := make(map[string]int)
  return m
}

func (m map[string]int) appendNewMap(key rune, value byte) {
  m[key] = value
}
