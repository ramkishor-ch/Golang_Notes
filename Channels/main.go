package main

import (
	"fmt"
	"net/http"
)

// Serial link checking, time consuming is more
func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	// Main Routine :
	// It will create child go routine
	// Main routine job is completed, nothing to run the code after creating child go routine.
	// Main routine will exit and from there will be no communication with "Child go routine".
	// To have a communication Main routine with child go routine then use channels.

	// Channels:
	// It communicates with each other different from go routine to go routine
	// or
	// from main routine to go routine.
	// After completion of execution of every child go routine then only main routine exits.

	c := make(chan string)

	for _, link := range links {
		// go routine : Enginer that executes code
		// go : Create a new thread go routine
		// checklink(link) : And run this function with it
		go checklink(link, c)
	}

	// Sending Data with Channels
	// channel <- 5 : Send the value '5' into this channel
	// myNumber <- channel : Wait for a value to be sent into the channel. When we get one,
	// 						 assign the value to 'myNumber'
	// fmt.Println(<-channel) : Wait for a value to be sent into the channel.
	//							When we get one log it out immediately.

	fmt.Printf(<-c)
}

func checklink(link string, c chan string) {
	// c : variable
	// chan : abbreviated type
	// string : value of the type string.

	// Depending call called Get(link) where it will wait for response
	// till then, "go routine" will sleep and it will tell to "main go rountine" that any other function would like to call
	// then go and call it.
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "might be down")
		c <- "Might be down I think"
		return
	}
	fmt.Println(link, "is working.")
	c <- "It is working"
}

//Parallel link checking, time consuming is less
//It can be achieved with Go routines

// Multiple Go routine will use Go Scheduler
// Go Scheduler : It runs one
// Scheduler runs one thread on EACH "logical" core
// by default Go tries to use one core!

// Concurrency :
// 1. We can have multiple threads executing code.
// 2. If one thread blocks, another one is picked up and worked on.
// 3. No need to wait for one Go routine to finsh before running next go rountine.
// 4. schedule work change between them on the fly

// Parallelism :
// 1. Multiple threads executed at the exact same time.
// Requires multiple CPU's
// 2. We can do multiple things at the same time.

// Running Program:
// Main Routine :
// 		It created when we launched our program
//
// Child go routine :
// 		It created by the 'go' keyword
