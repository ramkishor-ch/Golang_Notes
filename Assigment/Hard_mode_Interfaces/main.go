package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// fmt.Println(os.Args[1])

	// f, err := os.OpenFile("/Users/ramkishor.ch/Downloads/Fabric/ramkishor_Github_SoftLayer/Go_Local_/Assigment/Hard_mode_Interfaces/te.txt", os.O_CREATE, 0755)

	f, err := os.Open(os.Args[1])

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	// Reader interface
	// bs := make([]byte, 9999)
	// f.Read(bs)
	// fmt.Println(string(bs))

	io.Copy(os.Stdout, f)
}
