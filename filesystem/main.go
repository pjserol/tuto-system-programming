package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

var c struct {
	files int
	dirs  int
}

func main() {
	if len(os.Args) != 2 { // ensure path is specified
		panic("Specify a path!")
	}

	root, err := filepath.Abs(os.Args[1]) // get absolute path
	if err != nil {
		panic("cannot get absolute path!")
	}

	fmt.Printf("Listing files in %s:::\n\n", root)

	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		// walk the tree to count files and folders
		if info.IsDir() {
			c.dirs++
		} else {
			c.files++
		}
		fmt.Println("-", path)
		return nil
	})

	fmt.Printf("\nTotal: %d files in %d directories", c.files, c.dirs)

	// get extension of a file
	log.Printf("Extension:%s", filepath.Ext("test/main.go"))

	// get last element of the path
	log.Printf("Base:%s", filepath.Base("/Users/test/main.go"))
}
