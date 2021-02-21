package main

import (
	"fmt"
)

func main() {

	// option 1
	// Map with keys of type string and values of type string
	colors := map[string]string{
		"red":   "ff0000",
		"blue":  "00ff00",
		"green": "0000ff",
	}

	fmt.Println(colors)

	// option 2
	var animals map[string]string
	fmt.Println(animals)

	// option 3
	trees := make(map[string]string)
	trees["mahogany"] = "Zimbabwe"
	fmt.Println(trees)

}
