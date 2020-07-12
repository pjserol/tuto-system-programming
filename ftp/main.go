package main

import (
	"errors"
	"log"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		panic("Please specify a directory.")
	}

	path := os.Args[1]

	s, err := os.Stat(path)
	if err != nil {
		log.Fatalln("Invalid path:", err)
	}

	if err == nil && !s.IsDir() {
		err = errors.New("not a directory")
	}

	http.Handle("/", http.FileServer(http.Dir(path)))

	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal(err)
	}
}
