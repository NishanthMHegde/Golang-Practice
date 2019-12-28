package main

import (
	"fmt"
	"net/http"
)

func main() {
	websites := []string{
		"http://google.com",
		"http://facebook.com",
		"http://golang.org",
		"http://youtube.com",
		"http://cricbuzz.com",
	}
	for _, link := range websites {
		fmt.Println(link)
		//to create a Go routine, just add the prefix go in front of the function call
		//a Go sub routine executes the code lines within the function on a seperate thread.
		//Each go routine is executed one after the other until a blocking call (in our case, the http get request) is called.
		// Each go routine makes use of the same CPU core if only one core exists
		//If more than one core exists then each go routine makes use of one CPU core. THis is decided by Go CPU schedule.
		//Channels are used to communicate between one go routine and another go routine.
		//Channel needs to have a datatype and use the same datatype value across all routines (in our case it is string).
		c := make(chan string)
		//send the channel as parameter
		go checkLink(link, c)
		//read the channel contents
		fmt.Println(<-c)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println("Error", err)
		fmt.Println("The site is not up")
		// put the message into the channel
		c <- "The site is not up"
	}
	c <- "The site is up"
	fmt.Println("The site is up")

}
