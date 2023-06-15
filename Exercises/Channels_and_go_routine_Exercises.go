package main

// Question 1:
// Go Routines and Channels are tough, so let's start with the basics!
// Which of the following best describes what a go routine is?
// a. Code that is ran only when the go test command is used
// b. A seperate line of code execution that can be used to handle blocking code
// c. A seperate line of code execution that is *only* used for reading files from the
//    operating system
// Answer : b

// Question 2:
// What's the purpose of a channel?
// a. For communication between different functions in a single program
// b. For handling web socket requests
// c. For communication between go routines
// Answer : c

// Question 3:
// Take a look at the following program.  Are there any issues with it?
// package main

// import (
//  "fmt"
// )

// func main() {
//  greeting := "Hi There!"

//  go (func() {
//      fmt.Println(greeting)
//  })()
// }
// a. There are no issues
// b. The greeting variable is referenced from directly in the go routine, which
//    might lead to issues if we eventually start to change the value of greeting.
// c. The program will likely exit before the fmt.Println function has an opportunity
//    to actually print anything out to the terminal; this might not be the intent of the
//    program.
// d. Both answers #2 and #3 are correct
// Answer: d

// Question 4:
// Here's a tough one - Is there any issue with the following code?
// package main

// func main() {
//  c := make(chan string)
//  c <- []byte("Hi there!")
// }
// a. No issue, looks good! The channel is created to expect a value of type string,
//    but since we pass in a byte slice (which is like a string) we will be OK.
// b. The channel is expecting values of type string, but we are passing in a value
//    of type byte slice, which is not technically a string.
// Answer : b

// Question 5:
// Another tough one!  Is there any issue with the following code?
// package main

// func main() {
//      c := make(chan string)
//      c <- "Hi there!"
// }
// a. The syntax of this program is OK, but the program will never exit because it will
//    wait for something to receive the value we're passing into the channel.
// Explaination :
// fatal error: all goroutines are asleep - deadlock!
// goroutine 1 [chan send]:
// main.main()
//         /Users/ramkishor.ch/Downloads/Fabric/ramkishor_Github_SoftLayer/Go_Local_/Exercises/Channels_and_go_routine_Exercises.go:49 +0x34
// exit status 2
// b. There needs to be a comma between chan and string
// c. Looks good to me.
// Answer : a

// Question 6:
// Ignoring whether or not the program will exit correctly, are the following two code snippets equivalent?
// Snippet #1
// package main

// import "fmt"

// func main() {
//  c := make(chan string)
//  for i := 0; i < 4; i++ {
//      go printString("Hello there!", c)
//  }

//  for s := range c {
//      fmt.Println(s)
//  }
// }

// func printString(s string, c chan string) {
//  fmt.Println(s)
//  c <- "Done printing."
// }
// Snippet #2
// package main

// import "fmt"

// func main() {
//  c := make(chan string)

//  for i := 0; i < 4; i++ {
//      go printString("Hello there!", c)
//  }

//  for {
//      fmt.Println(<- c)
//  }
// }

// func printString(s string, c chan string) {
//  fmt.Println(s)
//  c <- "Done printing."
// }
// a. They are the same
// b. They will print out values in different orders
// c. #1 has syntax error, but #2 looks good
// d. #2 has a syntax error, but #1 looks good
// Answer : a
