package main

import "fmt"

// go Array == java Array
// go Slice == java Arraylist

func main() {
	// Slice
	cards := []string{"Ace of Diamonds", newCard(), newCard()}
	// Append returns a new slice
	cards = append(cards, "Six of spades")

	// := because we reinitialize the value of i and card with each iteration
	for i, card := range cards {
		fmt.Println(i, card)
	}

	fmt.Println(cards)
}

func newCard() string {
	return "Five of diamonds"
}
