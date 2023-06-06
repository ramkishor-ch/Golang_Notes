package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of 'deck'
// which is a slice of strings
type deck []string

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

// returning multiple values.
func deal(d deck, handsize int) (deck, deck) {
	return d[:handsize], d[handsize:]
}

// byte slice : Ascii of user input string
// []byte("Hi there!")
//[]byte: Type we want
//"Hi there!": Value we have
// Example:
// greeting := "Hi there!"
// fmt.Println([]byte(greeting))

// Deck to String and joining slice of strings
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

// save data to harddrive or file
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

// reading from the hard drive
func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	// bs: byte slice string of cards
	// err: value of type 'error', if nothing then value of 'nil'

	if err != nil {
		// Option #1 : log the error and return a call to newDeck()
		// Option #2 : log the error and entirely quit the program
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")
	return deck(s)
}

// Shuffling a Deck
// func (d deck) Shuffle() {
// 	for i := range d {
// 		newPosition := rand.Intn(len(d) - 1)

// 		d[i], d[newPosition] = d[newPosition], d[i]
// 	}
// }

// Random number generation
func (d deck) Shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
