package main

import "fmt"

type contactInfo struct {
	email    string
	postCode int
}

type person struct {
	// Properties
	firstName string
	lastName  string
	// Embedded Struct
	contact contactInfo
}

func main() {
	john := person{
		firstName: "John",
		lastName:  "Jones",
		contact: contactInfo{
			email:    "john@gmail.com",
			postCode: 4000,
		},
	}

	fmt.Printf("%+v \n", john)
}
