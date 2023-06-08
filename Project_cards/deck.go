package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// 5 Custom Type Declarations
// Create a new type of 'deck'
// which is a slice of strings
type deck []string

// 6 Receiver Functions
// receiver argument is 'd'
// The receiver is just a special case of a parameter.
// Go provides syntactic sugar "func (d deck) print() { }" to attach methods to types
// by declaring the first parameter as a receiver.

// deck: Every variable of type 'deck' can call this function on itself.
func (d deck) print() { // d is passed as reference to "cards_1" variables and
	// actual copy of the deck were working with is available in the function as a variable called 'd'
	for i, card := range d { // d: Python can be called as "self" or Ruby is "this" in other languages.
		fmt.Println(i, card)
	}
}

// 7 Creating New Deck
func newDeck() deck {
	cards := deck{}

	// declaring lists
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	// If you don't want to use index then you can mention it like "_"
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

// 8 Returning multiple values.
func deal(d deck, handsize int) (deck, deck) { // return two values of type deck
	return d[:handsize], d[handsize:]
}

// 10
// Deck to String and joining slice of strings
func (d deck) toString() string {
	// deck > []string > string > []byte
	return strings.Join([]string(d), ",")
}

// 11
// save data to harddrive or file
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

// 12
// reading from the hard drive
func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	// bs: byte slice string of cards
	// err: value of type 'error', if nothing then value of 'nil'

	// 13
	if err != nil {
		// Option #1 : log the error and return a call to newDeck()
		// Option #2 : log the error and entirely quit the program
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	// []byte > string > []string > deck
	s := strings.Split(string(bs), ",")
	return deck(s)
}

// 14
// Shuffling a Deck
func (d deck) Shuffle() {
	for i := range d {
		newPosition := rand.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}
}

// 15
// Random number generation
func (d deck) Shuffle_2() {
	source := rand.NewSource(time.Now().UnixNano()) // different int64 number every single time start of a program and we use it as a seed to generate a new source object.
	r := rand.New(source)                           // using that a new source object to generate a new number.

	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
