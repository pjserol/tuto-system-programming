package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		panic("Please specify a path.")
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}

	defer f.Close()

	// wrapping the reader with a buffered one
	r := bufio.NewReader(f)

	var rowCount int
	for err == nil {
		var b []byte
		for moar := true; err == nil && moar; {
			b, moar, err = r.ReadLine()
			if err == nil {
				fmt.Print(string(b))
			}
		}

		// each time moar is false, a line is completely read
		if err == nil {
			fmt.Println()
			rowCount++

		}
	}

	if err != nil && err != io.EOF {
		panic(err)
	}

	fmt.Println("\nRow count:", rowCount)
}
