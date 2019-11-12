// Introduction
// Calculate the number of grains of wheat on a chessboard given that the number on each square doubles.
//
// There once was a wise servant who saved the life of a prince. The king promised to pay whatever the servant could dream up. Knowing that the king loved chess, the servant told the king he would like to have grains of wheat. One grain on the first square of a chess board, with the number of grains doubling on each successive square.
//
// There are 64 squares on a chessboard (where square 1 has one grain, square 2 has two grains, and so on).
//
// Write code that shows:
//
// how many grains were on a given square, and
// the total number of grains on the chessboard

package main

import(
  "fmt"
  "math"
)


func grains_on_square(x float64) float64 {
  return math.Pow(2, x)
}

func all_grains(y float64) float64 {
  var sum float64
  for i := 0.0; i < y; i++ {
    value := grains_on_square(i)
		sum += value
	}
  return sum
}

func main() {
  // fmt.Println(grains_on_square(1))
  fmt.Println(all_grains(64))
}
