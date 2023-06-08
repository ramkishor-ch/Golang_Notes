package main

import "fmt"

func main() {
	// 1: Variable Declarations and Types
	// var: variable about to create new variable
	// card: name of the variable
	// string: string will be ever be assigned to this variable
	// Ace of Spades: assign the value
	// var card string = "Ace of Spades"
	// card_ := "Five of Diamonds"
	// fmt.Println(card, ",", card_)

	// card = "Four of Clubs"
	// fmt.Println(card)

	//Go, C++, Java: statically typed language
	//Javascript, Ruby, Python: dynamically typed language

	// Basic Go Types:
	// bool: true, false
	// string: "Hi"
	// int: 0, -10000, 99999
	// float64: 10.1, 0.00009, -199.01

	// 3 Array and Slice
	// Array: A Static array, fixed length list of things
	// Slice: A Dynamic array that can grow or shrink.
	// cards := []string{"Ace of Diamonds", newCard()}
	// cards = append(cards, "Six of Spades")
	// fmt.Println(cards)

	// 4 For Loops
	// i: index of this element in the array
	// card: Current card were iterating over
	// range cards: Take the slice of cards and loop over it
	// for i, card := range cards {
	// 	fmt.Println(i, card)
	// }

	// 5 Custom Type Declarations and 6 Receiver Functions
	// cards_1 := deck{"Ace of Diamonds", newCard()}
	// cards_1 = append(cards_1, "Six of Spades")
	// cards_1.print()

	// 7 Creating a New Deck
	// cards_2 := newDeck()
	// cards_2.print()

	// 8 Returning multiple values
	// cards_3 := newDeck()
	// hand, remainingCards := deal(cards_3, 5)
	// hand.print()
	// remainingCards.print()

	// 9 byte slice : Ascii of user input string
	// []byte("Hi there!")
	// []byte: Type we want
	// "Hi there!": Value we have
	// Example:
	// greeting := "Hi there!"
	// fmt.Println([]byte(greeting))

	// 10
	// Type conversion: the type we want
	// deck > []string > string > []byte
	// greeting := "welcome"
	// fmt.Println([]byte(greeting))
	// Deck is converted to []string is converted to String
	card1 := newDeck()
	fmt.Println(card1.toString())

	//11
	// string is converted to slice of bytes []byte
	// card2 := newDeck()
	// card2.saveToFile("my_cards") //my_cards is the filename

	// 12 and 13
	// filename
	// card2 := newDeckFromFile("my_cards")
	// card2.print()

	// 14
	// card3 := newDeck()
	// card3.Shuffle()
	// card3.print()

	// 15
	// card4 := newDeck()
	// card4.Shuffle_2()
	// card4.print()
}

// 2 Functions and Return Types
// func
// newCard() : define a function called 'newCard'
// string : return a value of type 'string' of the function
// func newCard() string {
// 	return "Five of Diamonds"
// }
