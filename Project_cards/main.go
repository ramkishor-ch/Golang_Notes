package main

import "fmt"

func main() {
	var card string = "Ace of Spades"
	// var: about to create new variable
	// card: name of the variable
	// string: string will be ever be assigned to this variable
	// Ace of Spades: assign the value

	//Go, C++, Java: statically typed
	//Javascript, Ruby, Python: dynamically typed

	// Basic Go Types:
	// bool: true, false
	// string: "Hi"
	// int: 0, -10000, 99999
	// float64: 10.1, 0.00009, -199.01

	card = newCard()
	fmt.Println(card)

	//Array: Fixed length list of things
	//Slice: An array that can grow or shrink.
	cards := deck{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of Spades")
	cards.print()

	// for i, card := range cards {
	// 	fmt.Println(i, card)
	// }
	//i: index of this element in the array
	//card: Current card were iterating over
	//range cards: Take the slice of cards and loop over it

}

// Functions and Return Types
func newCard() string {
	return "Five of Diamonds"
}

// func
// newCard() : define a function called 'newCard'
// string : return a value of type 'string' of the function
