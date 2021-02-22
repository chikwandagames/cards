package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// Custom writer
type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// Here we use the writer interface
	// here the Copy funtion takes data from resp.Body and writes it to the terminal
	// repo.Body implements the readar interface

	// io.Copy(os.Stdout, resp.Body)

	// Custom writer
	lw := logWriter{}
	io.Copy(lw, resp.Body)

}

// By just defining this Write function, and associating it with the logWriter struct
// The logWriter is now implementing the Writer interface
func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Printf("Just wrote %v bytes \n", len(bs))
	return len(bs), nil
}
