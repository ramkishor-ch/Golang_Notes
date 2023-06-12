package main

import "fmt"

func main() {
	// 1
	// What is map?
	// one way
	colors := map[string]string{
		"red":   "#ff0000",
		"white": "#45fae0",
		"blue":  "#fce045a",
	}

	// 2
	// Manipulating map
	// another way
	// colors := make(map[int]string)
	// colors[10] = "ffffff"

	// deleting the key 10
	// delete(colors, 10)
	// fmt.Println(colors)

	printmap(colors)
}

// 3
// Iterating through map
// c is argument name
// type : map[string]string
// declaration and initialization is completed at the same line: for key, value := range c{}
func printmap(c map[string]string) {
	for key, value := range c {
		fmt.Println("Hexacode for color", key, "is", value)
	}
}

// Difference between map and struct
//			 		Map								Struct
// 1. All key and values must be				1. Values can be of different type
//    same type.
// 2. Keys are indexed we can iterate			2. Keys don't support indexing
//	  over them
// 3. Reference type:							3. Value type
//    a. copying a reference to a map
//    b. passing a map of to a function
//       that is reference to the underline
//		 data structure.
// 4. Don't need all keys at compile			4. You need to know all the different fields at compile time.
//    time
// 5. Use to represent a collection of  		5. Use to represent a "thing" with a lot of different properties.
//    related properties
