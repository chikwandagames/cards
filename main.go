package main

import "fmt"

func main() {
	// Long form of varibale declaration
	var aStr string = "A random string"
	fmt.Println(aStr)

	// Short form, using the type of declaration
	// the compiler inferes the variable type
	// := is only used to define new variables
	cards := "Ace of Spaceds"
	fmt.Println(cards)
}
