package main

import (
	"fmt"
	"net/http"
)

// The function main creates a Main go routine by default,
// therefore all the other go routines created by go checkLink
// are child go routines
// the issue here is that the execution of main() completes before
// the child go routines have completed, which creates a bug
// to fix this we user CHANNELs

func main() {
	links := []string{
		"http://www.google.co.uk",
		"http://www.facebook.com",
		"http://www.stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	// CHANNEL
	// go routines can comminicate between each other by means of channels
	// Channels are typed, the data passed between routines must be of the
	// same type
	c := make(chan string)

	for _, link := range links {
		// go key word, makes this call a go routine, it prevents blocking!
		// The go schedullar will pass of execution to the next go routine
		// in this case another call to checkLink in the loop before completing
		// execution of the previous one
		go checkLink(link, c)
	}
	// fmt.Println(<-c), this receives data from a channel, and prints
	// Receiving messages for a channel is a BLOCKING operation

	for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
	}

}

// Takes a channel for communication between the go routines
func checkLink(link string, c chan string) {
	// http.Get() returns a struct representing the response and an error message
	// Here we are only interested in the error
	_, err := http.Get(link) // this is BLOCKING!!

	if err != nil {
		fmt.Println(link, "might be down!")
		// Send a message into the channel
		c <- "Might be down"
		return
	}
	fmt.Println(link, "is up!")
	// Send a message into the channel
	c <- "Yes it up"
}
