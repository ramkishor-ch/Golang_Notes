package main

import "fmt"

// Create a new type of 'deck'
// which is a slice of strings
type deck []string

// receiver argument is 'd'.
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// d: The actual copy of the deck were working with is available
// in the function as a variable called 'd'
// d: reference to cards variables
// d: can be called as self or this in other languages.

// deck: Every variable of type 'deck' can call this function on itself.
