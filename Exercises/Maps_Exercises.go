// Question : 1
// Take a look at the following code.  What would the print statement log out?
package main

import "fmt"

func main() {
	m := map[string]string{
		"dog": "bark",
	}

	changeMap(m)

	fmt.Println(m)
}

func changeMap(m map[string]string) {
	m["cat"] = "purr"
}

// a. Only map[dog: bark] because Go mage a copy of the map when it was passed to the function chnageMap
// b. map[gog: bark cat: purr]
// Answer: b

// Question : 2
// What would happen if we tried to run the following program?  Look closely at all the code in it :)
// package main
// import "fmt"

// func main() {
//  m := map[string]string{
//    "dog": "bark",
//    "cat": "purr",
//  }

//  for key, value := range m {
//    fmt.Println(value)
//  }
// }
// a. It would run successfully
// b. The compiler would throw an error because there is an extra comma after "cat": "purr"
// c. The compiler would throw an error because the variable key was created but never used.
// Answer: c
