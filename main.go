package main

import (
	"fmt"
)

func main() {
	// Map with keys of type string and values of type string
	colors := map[string]string{
		"red":   "ff0000",
		"blue":  "00ff00",
		"green": "0000ff",
		"white": "ffffff",
		"black": "000000",
	}
	colors["pink"] = "fff050"

	fmt.Println(colors)
	// Deleting a key and value
	delete(colors, "red")
	fmt.Println(colors)

	printMap(colors)
}

func printMap(cm map[string]string) {
	for k, v := range cm {
		fmt.Println(k + ": " + v)
	}
}
