// package dna
package main

import (
	"fmt"
)

// Histogram is a mapping from nucleotide to its count in given DNA.
// Choose a suitable data type.
type Histogram struct {
	A int
	C int
	G int
	T int
}

// DNA is a list of nucleotides. Choose a suitable data type.
type DNA struct {
	composition string
}

// Counts generates a histogram of valid nucleotides in the given DNA.
// Returns an error if d contains an invalid nucleotide.
///
// Counts is a method on the DNA type. A method is a function with a special receiver argument.
// The receiver appears in its own argument list between the func keyword and the method name.
// Here, the Counts method has a receiver of type DNA named d.
func (d DNA) Counts() (Histogram, error) {
	var h Histogram

	for index := 0; index < len(d.composition); index++ {
		if d.composition[index] == "A" {
			h.A++
		} else if d.composition[index] == "C" {
			h.C++
		} else if d.composition[index] == "G" {
			h.G++
		} else if d.composition[index] == "T" {
			h.T++
		}
	}
	return h, nil
}

func main() {
	var d DNA
	d.composition = "AACCGGT"
	fmt.Println(d.Counts())
}
