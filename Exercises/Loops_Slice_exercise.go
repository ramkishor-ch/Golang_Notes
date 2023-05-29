//Would the following code compile successfully?  Try it out yourself!

package main

import "fmt"

func main() {
	cards := []string{"Ace of Diamonds"}
	fmt.Println(cards)

	cards = append(cards, "Six of Spades")

	for index, card := range cards {
		fmt.Println(index, card)
	}
}
