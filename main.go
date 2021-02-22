package main

import "fmt"

// Interfaces are named collections of method signatures.

// type bot interface {},
// this means, any type that implements the method getGreeting is now of type
// bot as well
// therefore any value of type spanishBot or englishBot can now be passed to
// functions that accept type bot
// bot is of interface type, therefore we cannot directly create a value of bot
// spanishBot, englishBot, int, map, struct, string... are concrete types
type bot interface {
	getGreeting() string
}

type englishBot struct {
}

type spanishBot struct {
}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func (eb englishBot) getGreeting() string {
	/*
		Custom logic for genetation english greeting here
	*/
	return "Hi there!"
}

// I we are not using the spanish bot variable, we can leave it out completely
func (spanishBot) getGreeting() string {
	/*
		Custom logic for genetation spanish greeting here
	*/
	return "Holla!"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}
