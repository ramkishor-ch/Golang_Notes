package main

import "fmt"

type Englishbot struct{}
type Spanishbot struct{}

// bot : interface name
// Greeting() (): function name with arguments inside (string, int), (string, error) : list of return types.
type bot interface {
	Greeting() string
}

func main() {
	en := Englishbot{}
	sp := Spanishbot{}

	printGreeting(en)
	printGreeting(sp)
}

func printGreeting(b bot) {
	fmt.Println(b.Greeting())
}

// func printGreeting(en Englishbot) {
// 	fmt.Println(en.Greeting())
// }

// func printGreeting(sp Spanishbot) {
// 	fmt.Println(sp.Greeting())
// }

func (en Englishbot) Greeting() string {
	return "Welcome"
}

func (sp Spanishbot) Greeting() string {
	return "Wola"
}

// Interfaces :
// An interface type is defined as a "set of method signatures" but not their implementation.
// A value of interface type can hold any value that implements those methods.

// Why interfaces are required ?
// 1. We cannot define two same functions name with different types of having arguments, to avoid
// the problem we can use interfaces.
// 2. Reusability of code.
// 3. code optimization.
// 4. Every function we write has to be rewritten to accommodate different types even,
//    if the logic in it is identical.

// concrete Type : create a value each of them
// map, struct, int, string, englishBot

// Interface Type : we cannot a value directly of the interface type.
// Interfaces are not generic types
// Interfaces are implicit:
// 		example: where englishbot and spanishbot are belong to bot type.
// Interfaces are a contract to help us manage types.
//		If our custom types implementation of a function is broken then interfaces wont help us.
// Interfaces are tough.
//		Understanding how to read them in standard library
//		Writing your own interfaces is tough and requires experience.
