package main

import "fmt"

func main() {
	cards := []string{"Ace if Diamonds", newCard()}
	cards2 := deck{"Jack of Diamonds", newCard()}

	fmt.Println(cards)

	cards2.myPrint()
}

func newCard() string {
	return "Five of Diamonds"
}
