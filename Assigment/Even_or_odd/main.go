package main

import "fmt"

func main() {
	fmt.Println("Assignment")

	sliceoofints := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i, sints := range sliceoofints {
		if sints%2 == 0 {
			fmt.Println(i, "is even")
		} else {
			fmt.Println(i, "is odd")
		}
	}

}
