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

	// Pointers
	// johnPtr points to the address of john
	johnPtr := &john
	johnPtr.updateName("Johnny", "Bones")

	// Go allows you to call a function that receives ta pointer
	// with a pointer or actaual value
	johnPtr.myPrint()
	john.myPrint()

}

// receiver function
// * before type
// This is a type, it means we are woking with a pointer to a person
func (personPtr *person) myPrint() {
	fmt.Printf("%+v \n", (*personPtr))
}

func (personPtr *person) updateName(fName string, lName string) {
	// Dererence the pointer to person, change from pointer to value
	// * before value
	// This means we want to manipulate the value the pointer is referencing
	(*personPtr).firstName = fName
	(*personPtr).lastName = lName
}
