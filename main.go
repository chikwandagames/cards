package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://www.google.co.uk",
		"http://www.facebook.com",
		"http://www.stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	// go routines can comminicate between each other by means of channels
	// Channels are typed, the data passed between routines must be of the
	// same type

	for _, link := range links {
		// go key word, makes this call a go routine, it prevents blocking!
		// The go schedullar will pass of execution to the next go routine
		// in this case another call to checkLink in the loop before completing
		// execution of the previous one
		go checkLink(link)
	}
}

// This approach is very BAD, because its sequential, we check every link
// one by one then print
func checkLink(link string) {
	// http.Get() returns a struct representing the response and an error message
	// Here we are only interested in the error
	_, err := http.Get(link) // this is BLOCKING!!

	if err != nil {
		fmt.Println(link, "might be down!")
		return
	}
	fmt.Println(link, "is up!")
}

func makeRequestParallel() {

}
