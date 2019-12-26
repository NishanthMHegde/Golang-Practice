package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	//HTTP GET method which will get the HTML page of the desired URL entered otherwise an error.

	resp, err := http.Get("http://www.google.com")
	if err != nil {
		fmt.Println("Error:,", err)
		os.Exit(1)
	}

	//We will now use the Reader interface which is built inside an io.ReadWrite Interface (interface within an interface)
	//Reader interface accepts the input/response coming from outside (for ex: HTTP request's responses) and can be read into a byte slice

	bs := make([]byte, 999999) //create a byte slice of a big size which can hold the value read by the reader interface
	resp.Body.Read(bs)         //read the body of the response into the byte slice
	fmt.Println(string(bs))

	//Another method to do that would be to make use of the io.Copy
	//io.Copy(destination, source) takes an argument destination which is a Writer interface implementation and source which is a Reader interface implementation
	// The response read from the Reader interface is plugged into the Writer interface

	io.Copy(os.Stdout, resp.Body) //comment out the byte slice coe of 3 lines for this to work
	//The Writer interface is used to take input from program and send it outside
}
