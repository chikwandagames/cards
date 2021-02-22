package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// the byte slice for holding all the data added to it by the Read function
	// 999999 is the size or length
	bSlice := make([]byte, 99999)

	resp.Body.Read(bSlice)
	fmt.Println(string(bSlice))

}
