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

	// option 3
	var jane person
	// At this point firstName and lastName are zero values
	// Defult values e.g. int 0, bool false, string ""...
	fmt.Println(jane)
	// %+v, will print all the fields and their values
	fmt.Printf("%+v \n", jane)
	//
	jane.firstName = "Jane"
	jane.lastName = "Thomas"

	fmt.Println(alex)
	fmt.Println(bob)
}
