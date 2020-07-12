package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		panic("Please specify a path.")
	}

	filename := os.Args[1]

	// Read all the content at once from the memory
	readFileAtOnce(filename)

	fmt.Println("----------")

	readFileByChunks(filename)
}

func readFileAtOnce(filename string) {

	b, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))
}

func readFileByChunks(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	//  ensure to close the file to avoid leaks
	defer f.Close()

	b := make([]byte, 32)

	for n := 0; err == nil; {
		n, err = f.Read(b)
		if err == nil {
			fmt.Print(string(b[:n]))
			//fmt.Println("\n------------")
		}
		if err != nil && err != io.EOF {
			// expected an EOF at the end of the file
			fmt.Println("\n\nError:", err)
		}
	}
}
