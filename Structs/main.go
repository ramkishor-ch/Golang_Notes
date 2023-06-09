package main

import "fmt"

// 5 : Embedding Structs
type contactInfo struct {
	email   string
	zipCode int
}

// 1 : Structs in Go
// 2 : Defining Structs
type person struct {
	firstName string
	lastName  string
	contact   contactInfo // no need of contact but we can use contactInfo
}

func main() {
	// 3 : Declaring Structs
	// alex := person{"Alex", "Anderson"}
	// or
	// alex := person{firstName:"Alex", lastName:"Anderson"}

	// fmt.Println(alex)

	// 4 : Updating Struct Values
	// var alex person
	// fmt.Println(alex)
	// fmt.Printf("%+v \n", alex) // %+v: all field names and values of alex

	// alex.firstName = "Alex"
	// alex.lastName = "Anderson"
	// fmt.Println(alex)

	// 5
	jim := person{
		firstName: "jim",
		lastName:  "party",
		contact: contactInfo{ // contactInfo : contactInfo
			email:   "jim@gmail.com",
			zipCode: 521136,
		},
	}
	// fmt.Printf("%+v\n", jim)

	// 7 : Pass by value
	// jim.updateFirstName("jimmy")
	// jim.print()

	// 8, 9 : structs with pointers, Pointer Operations
	// &variable : Give me the memory address of the value this variable is pointing at
	// jimPointerHolder := &jim
	// jimPointerHolder.updateFirst_Name("jimmmy")
	// jimPointerHolder.print()

	// Go : provides automatically converts the type of person "jim" into pointer to a person
	// 10 : Pointer Shortcut
	// jim : type person
	jim.updateFirst_Name("gem")
	jim.print()

}

// Type  	Zero Value
// string	""
// int		0
// float	0
// bool		false

// 6
// Structs with receiver functions
// call this function of any value of any type this person
// inside body of the function you can reference the variable p
func (p person) print() {
	fmt.Printf("%+v", p)
}

// 7 : pass by value
func (p person) updateFirstName(newfirstName string) {
	p.firstName = newfirstName
}

// 8, 9 : structs with pointers, Pointer Operations
// *pointer : Giver me the value this memory address is poiniting at
// Receiver : *person is the type description
// *person : working with type of "pointer to a person"
// *pointerToPerson : we want to manipulate the value of pointer is referencing.
// *pointerToPerson is same as jimPointer
func (pointerToPerson *person) updateFirst_Name(newfirstName string) {
	// reference firstName property and update newfirstName
	(*pointerToPerson).firstName = newfirstName // or pointerToPerson.firstName = newfirstName
}

// Flow :
// address: 0001
// value: person{firstName: "jim"..}
// Turn "address" into "value" by only possible with "*address"
// Turn "value" into "address" by only possible with "&value"

// by default : pass by value

// Value Types : Use pointers to change these things in a function
// int
// float
// string
// bool
// structs

// Reference Types : Don't worry about pointers with these
// Slices
// maps
// channels
// pointers
// functions
