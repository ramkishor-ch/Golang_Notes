package main

// Question 1:
// Whenever you pass an integer, float, string, or struct into a function, what does Go do with that argument?
// a. It automatically creates a pointer to each argument
// b. It sets each argument to its 'zero value'
// c. It creates a copy of each argument, and these copies are used inside of the function
// Answer: c

// Question 2:
// import "fmt"

// func main() {
// 	name := "Bill"

// 	fmt.Println(&name)
// }
// Output : The memory address that "Bill" is stored at

// Question 3:
// What is the & operator used for?
// a. Joining two strings together
// b. Turning a value into a pointer
// c. Turning a pointer into a value
// Answer: Turning a value into a pointer

// Question 4:
// When you see a * operator in front of a pointer, what will it turn the pointer inot?
// a. A value
// b. it will remain a pointer
// c. A type definition
// Answer : a

// Question 5:
// When the following program runs,
// the fmt.Println call reports that the latitude field of newYork is still equal to 40.73 .
// What changes should we make to get the latitude of newYork to update to 41.0 ?
// package main
// import "fmt"

// type location struct {
// 	longitude float64
// 	latitude  float64
// }

// func main() {
// 	newYork := location{
// 		latitude:  40.73,
// 		longitude: -73.93,
// 	}

// 	newYork.changeLatitude()

// 	fmt.Println(newYork)
// }

// func (lo location) changeLatitude() {
// 	lo.latitude = 41.0
// }

// a. This is not possible with go, we cannot change the original value of a struct in a function.
// b. The changeLatitude function should use lo["latitude"] = 41.0 to update the struct
// c. Change the receiver type of changeLatitude to *location, the replace lo with
//    *lo in the function body. This will turn the pointer lo into a value type and then update it.

// Answer: c

// Question 6:
// package main

// import "fmt"

// type location struct {
// 	longitude float64
// 	latitude  float64
// }

// func main() {
// 	newYork := location{
// 		latitude:  40.73,
// 		longitude: -73.93,
// 	}

// 	newYork.changeLatitude()

// 	fmt.Println(newYork)
// }

// func (lo *location) changeLatitude() {
// 	(*lo).latitude = 41.0
// }

// In the 'changeLatitude' function, what is *location in the receiver list (after the word 'func') communicating to us?
// a. It specifies the type of the receiver that the function expects.
// b. It is trying to turn the location type memory address into a value.
// Answer : a

// Question 7
// Take a look at the following program.  What will the Println  function in the main  function print out?
// package main

// import "fmt"

// func main() {
// 	name := "Bill"
// 	updateValue(name)
// 	fmt.Println(name)
// }

// func updateValue(n string) {
// 	n = "Alex"
// }

// a. Bill
// b. "Alex"
// c. An empty string
// Answer : "Bill"

// Question 8:
// Take a look at the following program.  The changeLatitude function expects a receiver of type pointer to a location struct , but in the main function the receiver is a value type of a struct.  What will happen when this code is executed?
// package main

// import "fmt"

// type location struct {
// 	longitude float64
// 	latitude  float64
// }

// func main() {
// 	newYork := location{
// 		latitude:  40.73,
// 		longitude: -73.93,
// 	}

// 	newYork.changeLatitude()

// 	fmt.Println(newYork)
// }

// func (lo *location) changeLatitude() {
// 	(*lo).latitude = 41.0
// }

// a. It will crash
// b. The program will run, but it will make a copy of the newYork struct in The
//    changeLatitude function, so newYork will not be changed
// c. This program uses a shortcut, where Go will automatically assume that even
//    though newYork.changeLatitude() is using a value type we probably mean to pass
//    in a pointer to the newYork struct.
// Answer: c

// Question 9:
// Here's a tricky one!  What will the following program print out?
// package main

// import "fmt"

// func main() {
//     name := "Bill"

//     fmt.Println(*&name)
// }
// a. The string "Bill"
// b. The memory address that "Bill" is stored at
// c. The program will throw an error
// Answer: a
