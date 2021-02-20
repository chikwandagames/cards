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
	contactInfo
}

func main() {
	john := person{
		firstName: "John",
		lastName:  "Jones",
		contactInfo: contactInfo{
			email:    "john@gmail.com",
			postCode: 4000,
		},
	}

	john.updateName("Johnny", "Bones")
	john.myPrint()

}

// receiver function
func (p person) myPrint() {
	fmt.Printf("%+v \n", p)
}

func (p person) updateName(fName string, lName string) {
	p.firstName = fName
	p.lastName = lName
}
