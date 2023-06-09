package main

import "fmt"

// struct :
// Data structure, collection of properties that are related together.

func main() {
	myslice := []string{"Hi", "There", "how", "are", "you"}

	updateSlice(myslice)

	fmt.Println(myslice)
}

func updateSlice(s []string) {
	s[0] = "Bye"
}
