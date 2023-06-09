package main

// Question 1:
// When we create a slice, Go will automatically create which two data structures?
// a. An array and a string
// b. An array and a list of pointers
// c. An array and a structure that records the length of the slice,
// the capacity of the slice, and a reference to the underlying array
// Answer: c

// Question 2:
// In the following code snippet, when we pass mySlice to the updateSlice function, is the mySlice value being copied before being passed into the function?
// package main

// import "fmt"

// func main() {
//  mySlice := []string{"Hi", "There", "how", "are", "you?"}

//  updateSlice(mySlice)

//  fmt.Println(mySlice)
// }

// func updateSlice(s []string) {
//  s[0] = "Bye"
// }
// a. Yes
// b. no
// Answer: a

// Question 3:
// With 'value types' in Go, do we have to worry about pointers if we want to pass a value to a function and modify the original value inside the function?
// a. Yes
// b. no
// Answer: a

// Question 4:
// With 'reference types' in Go, do we have to worry about pointers if we want to pass a value to a function and modify the original value inside the function?
// a. Yes
// b. no
// Answer: b

// Question 5:
// Is a slice a 'value type' or a 'reference type'
// Answer:
// Reference type, because a slice contains a reference to the actual underlying list of records.

// Question 6:
// Here's a tough one.
// We've been discussing about how to use pointers to avoid passing something to a function by value.  So instead of passing a value to a function, we pass a pointer to that value.  But we've also said many times that Go is a "pass by value" language, which *always* copies arguments that are passed to a function.  Take a glance at the following code snippet, which gets a pointer to name called namePointer and prints out the memory address of the pointer itself!
// Then in the function, we take the pointer that was passed as an argument and print out the address of that pointer too.
// Do you think the memory address printed by both Println calls will be the same?  Why or why not?

// package main

// import "fmt"

// func main() {
//  name := "bill"

//  namePointer := &name

//  fmt.Println(&namePointer)
//  printPointer(namePointer)
// }

// func printPointer(namePointer *string) {
//  fmt.Println(&namePointer)
// }
// a. It will be the same because pointers aren't copied when passed to a function.
// b. The log statements will print different addresses because *everything* in Go is pass by value.
// Answer: b
