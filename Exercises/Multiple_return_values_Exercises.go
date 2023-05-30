// 1
package main

import "fmt"

func main() {
	color1, color2, color3 := colors()

	fmt.Println(color1, color2, color3)
}

func colors() (string, string, string) {
	return "red", "yellow", "blue"
}

//Output: "red" "yellow" "blue"

// 2
// package main

// import "fmt"

// func main() {
//    c := color("Red")

//    fmt.Println(c.describe("is an awesome color"))
// }

// type color string

// func (c color) describe(description string) (string) {
//    return string(c) + " " + description
// }

// Output: "Red is an awesome color"
// 'describe' is a function with a receiver of type 'color' that requires an argument
// of type 'string', then returns a value of type 'string'

// 3
// This is a tricky question about something that we'll cover in much greater detail later on.
// Think back to our current "cards" program, where we have the following code.
// func main() {
//     cards := newDeck()

//     hand, remainingCards := deal(cards, 5)

//     hand.print()
//     remainingCards.print()
// }
// After calling "deal" and passing in "cards", does the list of strings that the "cards" variable point at change?
// In other words, did we modify the 'cards' slice by calling 'deal'?

// We created two new references that point at subsections of the 'cards' slice.
// We never directly modified the slice that 'cards' is pointing at.
