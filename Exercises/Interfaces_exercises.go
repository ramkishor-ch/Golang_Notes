package main

// Question 1:
// When we say that interfaces can be satisfied implicitly, we mean that..
// a. We don't have to write any extra code to say that some type satisfies an interface
// b. We need to write code that declares custom types satisfy an interface.
// Answer: a

// Question 2:
// To say that a type satisfies an interface means that...
// a. The type has the same name as the interface
// b. The type implements all of the functions contains in the interface definition.
// c. The type and the interface both are used in the same package.
// Answer : b

// Question 3:
// In the following code, does the square type satisfy the shape interface?
// type shape interface {
//     area() int
// }

// type square struct {
//     sideLength int
// }

// func (s square) area() int {
//     return s.sideLength * s.sideLength
// }

// func printArea(s shape) {
//     fmt.Println(s.area())
// }
// a. Yes, because square defines the area function that returns an interface
// b. Yes, because the printArea function can be called with anything of the type shape
// c. No
// Answer: a

// Question 4:
// Take a look at the following code.  Does the rectangle type satisfy the shape  interface?
// type shape interface {
//     area() int
// }

// type square struct {
//     sideLength int
// }

// type rectangle struct {
//     height float64
//     width float64
// }

// func (s square) area() int {
//     return s.sideLength * s.sideLength
// }

// func (r rectangle) area() float64 {
//     return r.height * r.width
// }

// func printArea(s shape) {
//     fmt.Println(s.area())
// }
// a. Yes, because it defines the are function
// b. Yes, because it has the same method names as type square
// c. No, because rectangle's version of the area function returns a float64,
//    but the shape interface expects a return type of int
// Answer: c

// Question 5:
// Take a look at the following code.  Type square appears to successfully implement the shape interface, but the implementation of square 's area function looks broken - it always returns a value of 10 no matter what the side length of the square is.  Will the shape  interface do anything to help us catch this error?
// type shape interface {
//     area() int
// }

// type square struct {
//     sideLength int
// }

// func (s square) area() int {
//     return 10
// }

// func printArea(s shape) {
//     fmt.Println(s.area())
// }
// a. Yes, the interface should detect that we aren't doing our geometry correctly.
// b. Yes, the program should throw an error when we try to compile it.
// c. No, interfaces are only used to help with types. We can still easily write code
//    that does something completely wrong.
// Answer: c

// Question 6:
// Types that implement the Reader interface are generally used to...
// a. Read information from an outside data source into our application.
// b. Pipe data from one struct to another
// c. Read information from our program to an outside data source
// Answer: a

// Question 7:
// Imagine that you ask a coworker to create a new type that implements the Reader  interface to take data from a text file and print it on the screen.  They present you with the following code:
// type textFileReader struct {}

// func (textFileReader) Read(bs []byte) (int, error) {
//     return "Information from a text file"
// }
// They say that this code successfully compiled, so it must be correct. You then review their code and give them feedback.  What would you say?
// a. I would say that while the textFileReader type confirms to the requirements of the Reader
// interface, it doesn't actually implement the desired behaviour of reading a file from the
// hard drive
// b. I would say "nice job!"
// Answer: a

// Question 8:
// Because interfaces are satisfied implicitly, it can be tough to figure out which types satisfy which interfaces.
// Take a look at the type File inside the os package here: https://golang.org/pkg/os/#File
// Does the File type satisfy both the Reader and Writer interfaces?
// a. Yes
// b. No
// Answer: a
