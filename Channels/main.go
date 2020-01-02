package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	websites := []string{
		"http://google.com",
		"http://facebook.com",
		"http://golang.org",
		"http://youtube.com",
		"http://cricbuzz.com",
	}
	//to create a Go routine, just add the prefix go in front of the function call
	//a Go sub routine executes the code lines within the function on a seperate thread.
	//Each go routine is executed one after the other until a blocking call (in our case, the http get request) is called.
	// Each go routine makes use of the same CPU core if only one core exists
	//If more than one core exists then each go routine makes use of one CPU core. THis is decided by Go CPU schedule.
	//Channels are used to communicate between one go routine and another go routine.
	//Channel needs to have a datatype and use the same datatype value across all routines (in our case it is string).
	c := make(chan string)
	//send the channel as parameter
	for _, link := range websites {
		go checkLink(link, c)

	}
	//read the channel contents
	//reaing the channel is also a blocking call.
	// fmt.Println(<-c) //the main routine itself terminates after this call as it receives an answer from the blocking call

	//to prevent the above case of main routine terminating after reading one channel content, we can loop through  times and read output

	// for i := 0; i < len(websites); i++ {
	// 	fmt.Println(<-c)
	// }

	//But we need a progam which keeps checking the status of the websites without stop.
	//To achieve this, we need to call/create a sub routine everytime a websites status is checked and the channel is populated with the website ink from a previous routine

	//this will run permanentely and subroutine is called whenever a website is checked and channel has a value populated
	for link := range c {
		//We need to add a delay or sleep of 5 seconds between every time we check the status of a webiste.
		//We cannot call sleep directly in the function as it will result in the main routine going to sleep and blocks all other routines

		//Therefore, we need to create a function literal( AKA lambda function or anonymous function)
		//We also need to pass the link we have obtained to the function literal as a query parameter.
		//This is because every time the child routine goes to sleep, the link value is incremented in the main routine and another value is filled in the same location where the original pointer was pointing
		go func(link string) {

			time.Sleep(5 * time.Second)
			go checkLink(link, c)
		}(link) // This is where we call the anonymous function. we pass the link parameter as an argument which we obtained from the loop.
	}
}

func checkLink(link string, c chan string) {
	// fmt.Println(link)
	//Do not sleep here as it will lead to all chld routines going to sleep in the beginning itself
	//time.Sleep(5 * time.Second)

	_, err := http.Get(link)
	if err != nil {
		fmt.Println("Error", err)
		fmt.Println("The site ", link, " is not up")
		// put the message into the channel
		c <- link
		return
	}
	fmt.Println("The site ", link, "is up")
	c <- link

}
