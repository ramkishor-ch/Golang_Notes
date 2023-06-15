package main

import (
	"fmt"
	"net/http"
	"time"
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

	// Blocking Channel :
	// fmt.Printf(<-c) // only http://google.com is working because the other child go routine
	// are waiting for receive the message.
	// Once main routine received the message and then print it and main routine will exit it,
	// because main routine has only received message from http://google.com only after that
	// not received message from other facebook, stackoverflow tc.

	// Receive Messages :
	// for i := 0; i < len(links); i++ {
	// 	fmt.Printf(<-c)
	// }

	// Repeating Routines :
	// If I want to run same google.com for the second time then we use repeating routines.
	// The number of messages is equal to number links we make.
	//
	// for i := 0; i < len(links); i++ {
	// 	go checklink(<-c, c)
	// }
	// for {
	// 	go checklink(<-c, c)
	// }

	// Infinite loop
	// for {
	// }

	// Alternative Loop Syntax
	// same as infinite loop but this loop provides better readbility
	// for l := range c {
	// 	go checklink(l, c)
	// }

	// Sleeping Routine
	// for l := range c {
	// 	time.Sleep(7 * time.Second) // Duration type of int64
	// 	// Main Routine will pause for 7 seconds before next iteration of Main Routine.
	// 	go checklink(l, c)
	// }

	// Function Literals == Lambda Function == Anonymous Function
	// for l := range c {
	// 	go func() {
	// 		time.Sleep(7 * time.Second) // Duration type of int64
	// 		// Main Routine will pause for 7 seconds before next iteration of Main Routine.
	// 		checklink(l, c)
	// 	}()
	// }

	// Channels Gotcha
	// never ever share varibles directly between different child routine
	// because a same variable from different child routine point out at same address.
	//
	// Recommended :
	// Ever try to access the same variable from a different child routine wherever possible.
	// We only share information with a child routine or a new go routine that we create by passing it in as
	// an "argument" or communicating with the child routine over "channels".
	for l := range c {
		go func(l string) {
			time.Sleep(7 * time.Second) // Duration type of int64
			// Main Routine will pause for 7 seconds before next iteration of Main Routine.
			checklink(l, c)
		}(l)
	}
}

func checklink(link string, c chan string) {
	// purpose: checklink right now don't pause.

	// In checklink() go child routine() will pause but it is not recommended
	// to put it here because of the purpose so don't put time.Sleep in checklink()
	// time.Sleep(7 * time.Second)

	// c : variable
	// chan : abbreviated type
	// string : value of the type string.

	// Depending call called Get(link) where it will wait for response
	// till then, "go routine" will sleep and it will tell to "main go rountine" that any other function would like to call
	// then go and call it.
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "might be down")
		// c <- "Might be down I think"
		c <- link
		return
	}
	fmt.Println(link, "is working.")
	// c <- "It is working"
	c <- link
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
