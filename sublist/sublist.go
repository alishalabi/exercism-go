// Introduction
// Given two lists determine if the first list is contained within the second list, if the second list is contained within the first list, if both lists are contained within each other or if none of these are true.
//
// Specifically, a list A is a sublist of list B if by dropping 0 or more elements from the front of B and 0 or more elements from the back of B you get a list that's completely equal to A.
//
// Examples:
//
// A = [1, 2, 3], B = [1, 2, 3, 4, 5], A is a sublist of B
// A = [3, 4, 5], B = [1, 2, 3, 4, 5], A is a sublist of B
// A = [3, 4], B = [1, 2, 3, 4, 5], A is a sublist of B
// A = [1, 2, 3], B = [1, 2, 3], A is equal to B
// A = [1, 2, 3, 4, 5], B = [2, 3, 4], A is a superlist of B
// A = [1, 2, 4], B = [1, 2, 3, 4, 5], A is not a superlist of, sublist of or equal to B



package main

import (
  "fmt"
)


func isSubset(list1 []int, list2 []int) bool {
  contains := false
  // Base case: ensure list1 and list2 both have items
  if len(list1) == 0 || len(list2) == 0 {
    return false
  }
  // Iterate through list2, see if first item in list1 is contained
  for i := 0; i < len(list2); i++ {
    // If contained, see if each item list1 follows (in order)
    if list1[0] == list2[i] {
      lengthDifference := len(list2) - len(list1)
      if list1 == list2[i:lengthDifference] {
        contains := true
      }
    }
  }
  return contains
}

func main() {
  a := []int{3, 4}
  b := []int{1, 2, 3, 4, 5}

  twoContainsOne := isSubset(a, b)
  oneContainsTwo := isSubset(b, a)

  if twoContainsOne == true && oneContainsTwo == true {
    fmt.Println("A is equal to B")
  } else if twoContainsOne == true {
    fmt.Println("A is a subset of B")
  } else if oneContainsTwo == true {
    fmt.Println("B is a subset of A")
  }
}
