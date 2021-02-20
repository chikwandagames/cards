package main

import "fmt"

type person struct {
	// Properties
	firstName string
	lastName  string
}

func main() {
	// Declration
	// option 1
	alex := person{"Alex", "Anderson"}

	// option 2
	bob := person{firstName: "Bob", lastName: "Marley"}

	// option 2

	fmt.Println(alex)
	fmt.Println(bob)
}
