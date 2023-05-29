// In main.go: Files in the same package can freely call functions defined in other files.

package main

func main() {
	printState()
}

// In a separate file called state.go:
// package main

// import "fmt"

func printState() {
	// fmt.Println("California")
}
