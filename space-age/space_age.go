// Introduction
// Given an age in seconds, calculate how old someone would be on:
//
// Mercury: orbital period 0.2408467 Earth years
// Venus: orbital period 0.61519726 Earth years
// Earth: orbital period 1.0 Earth years, 365.25 Earth days, or 31557600 seconds
// Mars: orbital period 1.8808158 Earth years
// Jupiter: orbital period 11.862615 Earth years
// Saturn: orbital period 29.447498 Earth years
// Uranus: orbital period 84.016846 Earth years
// Neptune: orbital period 164.79132 Earth years
// So if you were told someone were 1,000,000,000 seconds old, you should be able to say that they're 31.69 Earth-years old.
//
// If you're wondering why Pluto didn't make the cut, go watch this youtube video.


package main

import (
  "fmt"
)

func convertAge(age float64, planet string) float64 {
  newAge := age
  if planet == "Mercury" {
    newAge = age / 0.2408467
  } else if planet == "Venus" {
    newAge = age / 0.61519726
  } else if planet == "Earth" {
    newAge = age / 1.0
  } else if planet == "Mars" {
    newAge = age / 1.8808158
  } else if planet == "Jupiter" {
    newAge = age / 11.862615
  } else if planet == "Saturn" {
    newAge = age / 29.447498
  } else if planet == "Uranus" {
    newAge = age / 84.016846
  } else if planet == "Neptune" {
    newAge = age / 164.79132
  }
  return newAge
}

func main() {
  fmt.Println(convertAge(10000000.0, "Mercury"))
  fmt.Println(convertAge(10000000.0, "Venus"))
  fmt.Println(convertAge(10000000.0, "Earth"))
  fmt.Println(convertAge(10000000.0, "Mars"))
  fmt.Println(convertAge(10000000.0, "Jupiter"))
  fmt.Println(convertAge(10000000.0, "Saturn"))
  fmt.Println(convertAge(10000000.0, "Uranus"))
  fmt.Println(convertAge(10000000.0, "Neptune"))
}
