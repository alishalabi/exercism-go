"""
Introduction:
Given a moment, determine the moment that would be after a gigasecond has passed.
A gigasecond is 10^9 (1,000,000,000) seconds.
"""


// Package gigasecond should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package gigasecond

// import path for the time package from the standard library
import (
	"time"
	// "fmt"
	)

// AddGigasecond should have a comment documenting it.
func AddGigasecond(t time.Time) time.Time {
	afterGigaSecond := t.Add(time.Second * 1000000000)
	return afterGigaSecond
}
