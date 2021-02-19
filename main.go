package main

import "fmt"

func main() {

	cards := newDeck()
	// cards.myPrint()

	// because neel returns 2 items
	hand, remainingCards := deal(cards, 5)
	hand.myPrint()
	fmt.Println("")
	remainingCards.myPrint()
}

func newCard() string {
	return "Five of Diamonds"
}
