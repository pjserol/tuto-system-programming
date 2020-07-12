package main

import (
	"fmt"
	"io"
	"os"
)

func exitCmd(w io.Writer, args ...string) bool {
	fmt.Fprintf(w, "Exit!\n")
	return true
}

func helpCmd(w io.Writer, args ...string) bool {
	fmt.Fprintln(w, "Available commands:")
	for _, c := range cmds {
		fmt.Fprintf(w, "  - %-15s %s\n", c.Name, c.Help)
	}
	return false
}

func printCmd(w io.Writer, args ...string) bool {
	if len(args) != 1 {
		fmt.Fprintln(w, "Please specify a file!")
		return false
	}

	f, err := os.Open(args[0])
	if err != nil {
		fmt.Fprintf(w, "Cannot open %s: %s\n", args[0], err)
	}

	defer f.Close()

	if _, err := io.Copy(w, f); err != nil {
		fmt.Fprintf(w, "Cannot print %s: %s\n", args[0], err)
	}

	fmt.Fprintln(w)

	return false
}
