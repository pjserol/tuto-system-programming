package main

import (
	"bytes"
	"fmt"
	"os"
)

const author1 = "Albert Camus"

type book struct {
	Author, Title string
	Year          int
}

func main() {
	filename := "book.txt"

	dst, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}

	defer dst.Close()

	bookList := []book{
		{Author: author1, Title: "The Stranger", Year: 1942},
		{Author: author1, Title: "The Plague", Year: 1947},
		{Author: author1, Title: "The Fall", Year: 1956},
	}

	b := bytes.NewBuffer(make([]byte, 0, 16))
	for _, v := range bookList {
		fmt.Fprintf(b, "%s - %s", v.Title, v.Author)
		if v.Year > 0 {
			fmt.Fprintf(b, " (%d)", v.Year)
		}

		b.WriteRune('\n')

		if _, err := b.WriteTo(dst); err != nil {
			fmt.Println("Error:", err)
			return
		}
	}

	fInfo, err := os.Stat(filename)
	if err != nil {
		panic(err)
	}

	fmt.Printf("File info:\n%+v", fInfo)
}
