package main

import "fmt"

// 1
type book string

func (b book) printTitle() {
	fmt.Println(b)
}

func main() {
	var b book = "Harry Potter"
	b.printTitle()
}

//Output: "Harry Potter"

// 2
// By creating a new type with a function that has a receiver, we...
// a. Are adding a 'method' to any value of that type
// b. Can call the function with any type we want

// Answer: a

// 3
// In the following snippet, what does the variable 'ls' represent?
// type laptopSize float

// func (ls laptopSize) getSizeOfLaptop() {
//     return ls
// }
// a. A value of type 'laptopSize'
// b. The type 'laptopSize'
// c. A function

// Answer: a

// 4
// Is the following code valid?
// type laptopSize float64

// func (this laptopSize) getSizeOfLaptop() laptopSize {
//     return this
// }

// a. No, it would fail to compile
// b. Yes, and there is nothing wrong with it.
// c. Yes, but it is breaking convention, Go avoids any mention of 'this' or 'self'

// Answer: c
