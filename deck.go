package main

// "string" package lets you use the join function
import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// A custom type extends a base type and adds some funtionality to it

// Create an new type deck, which is a slice of strings
// type deck == []string
type deck []string

// (d deck) is a receiver, i.e.
// Any variable of type deck now has access to myPrint method
// d is the current deck item we are working with, remember d is []string

// Calling myPrint on an item of type deck, will pass that item as a receiver
// e.g myPrint(item) then print it using the for loop

func (d deck) myPrint() {
	for i, card := range d { // d is like self or this
		fmt.Println(i, card)
	}
}

func newDeck() deck {
	cards := deck{}
	suits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	values := []string{"Ace", "Two", "Three", "Four", "Five"}

	for _, suit := range suits {
		for _, value := range values {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// (deck, deck), means we want to return two values of type deck
func deal(d deck, handSize int) (deck, deck) {
	// Slicing
	return d[:handSize], d[handSize:]
}

// A byte slice is a string repesented in ascii
// eg "Hi there!" == [72, 105, 32, 116, 104, 101, 114, 101, 33]

// Converts a deck into a string
func (d deck) toString() string {
	// Join concatenates all the elements present in the slice of string into a single string
	return strings.Join([]string(d), ", ")
}

func (d deck) saveToFile(filename string) error {
	// for saving a byte slice to file, we need to convert deck to string
	// then string to byte slice
	// 0666 permissions means anyone can read and write to the file
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	// ReadFile returns 2 items
	bSlice, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println("Error: ", err)
		// Exit the program, code other than 0, means somthing went wrong
		os.Exit(1)

	}

	// Cast the byte slice to string
	// Split func, splits a string into all substrings separated by the given separator
	// and returns a slice which contains these substrings
	ss := strings.Split(string(bSlice), ", ")
	return deck(ss)

}

func (d deck) shuffle() {
	// so that we have random seed each time we loop
	// we use time.Now to generate a unique int64 number each time we call this shuffle
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	// The item can be left out
	for i := range d {
		// random number from 0 to d length -1
		newPos := r.Intn(len(d) - 1)

		// Swap
		d[i], d[newPos] = d[newPos], d[i]
	}
}
