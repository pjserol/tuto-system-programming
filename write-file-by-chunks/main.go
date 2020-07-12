package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		panic("Please specify a path and some content!")
	}

	src, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}

	defer src.Close()

	// OpenFile allows to open a file with any permissions
	// O_APPEND append at the end of the destination
	dst, err := os.OpenFile(os.Args[2], os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}

	defer dst.Close()

	// go to the end of the source file
	cur, err := src.Seek(0, io.SeekEnd)
	if err != nil {
		panic(err)
	}

	b := make([]byte, 16)
	for step, r, w := int64(16), 0, 0; cur != 0; {

		// ensure cursor is 0 at max
		if cur < step {
			b, step = b[:cur], cur
		}

		cur = cur - step

		// go backwards on the source file
		_, err = src.Seek(cur, os.SEEK_SET)
		if err != nil {
			break
		}

		// read small part of the file
		if r, err = src.Read(b); err != nil || r != len(b) {
			// all buffer should be read
			if err == nil {
				err = fmt.Errorf("read: expected %d bytes, got %d", len(b), r)
			}
			break
		}

		// reverse the content
		for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
			switch { // Swap (\r\n) so they get back in place
			case b[i] == '\r' && b[i+1] == '\n':
				b[i], b[i+1] = b[i+1], b[i]
			case j != len(b)-1 && b[j-1] == '\r' && b[j] == '\n':
				b[j], b[j-1] = b[j-1], b[j]
			}

			// swap bytes
			b[i], b[j] = b[j], b[i]
		}

		// write to the destination
		if w, err = dst.Write(b); err != nil || w != len(b) {
			if err != nil {
				err = fmt.Errorf("write: expected %d bytes, got %d", len(b), w)
			}
		}
	}

	// expect an EOF
	if err != nil && err != io.EOF {
		fmt.Println("\n\nError:", err)
	}
}
