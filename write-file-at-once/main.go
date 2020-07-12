package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		panic("Please specify a path and some content!")
	}

	// needs to be casted to a byte slice
	data := []byte(os.Args[2])

	// writes all the content at once in a single operation
	if err := ioutil.WriteFile(os.Args[1], data, 0644); err != nil {
		fmt.Println("Error:", err)
	}
}
