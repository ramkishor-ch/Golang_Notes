// Print "Welcome to Go!" message

package main //package is a collection of common source code files

import "fmt" // automatically format all codes in each file in the current directory.

func main() {
	fmt.Println("Welcome to Go!") //
}

// go is not a oop language
// we don't use class terms etc.

//go run file command: will compile and execute it.
//command: go run filename.go

//go build: it will only compile the file will not execute it.

//go install: compiles and installs a package

//go get: downloads the raw source code of someone else's package.

//go test: run any tests files associated with the current project.

//Packages:

//Types:
//1. Executable: Generates a file that we can run
// Example: main
//2. Reusable: Code used as 'helpers' Good place to put reusable logic

//flow:
//Package main > go build > main.exe
//main.exe: if we ran this file, the function named 'main' would be automatically ran.

//Executable package:
//package main: "main" is special
//Defines a package that can be compiled and then "executed".
//Must have a function called "main".

//Reusable package:
//package calculator:
//Defines a package that can be used as a dependency: helper code.

//Import Statements:

// fmt: standard library package
// fmt: format

// fmt mean?
// access to other libraries like debug, math, encoding, io etc.

//more information: golang.org/pkg

// File Organization:

// package main : package declaration

// import "fmt" : Import other packages that we need

//		func main() { }
// func: declare a function
// main: sets the name of the function
// () list of arguments to pass the function
// {} function body: calling function runs this code.

// Reference: https://github.com/StephenGrider/GoCasts/tree/master/diagrams
