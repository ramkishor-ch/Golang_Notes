package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("https://google.com")

	lg := logWriter{}

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	// Reader interface
	// bs := make([]byte, 9999)
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))

	// Writer interface: take info and send it outside of our program
	// We need to find something in the standard library that *implements" the writer interface,
	// and use that to log out all the data that we are receiving from the Reader.
	// io.Copy(os.Stdout, resp.Body)

	// Custom Writer
	io.Copy(lg, resp.Body)
}

// Custom Writer
func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Bytes are returned: ", len(bs))
	return len(bs), nil
}
