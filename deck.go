package main

import "fmt"

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
