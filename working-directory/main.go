package main

import (
	"fmt"
	"os"
)

func main() {
	// get current directory
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("starting dir:", wd)

	// change directory
	if err := os.Chdir("../."); err != nil {
		fmt.Println(err)
		return
	}

	if wd, err = os.Getwd(); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("final dir:", wd)

}
