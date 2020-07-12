package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type cmd struct {
	Name   string // command name
	Help   string // description
	Action func(w io.Writer, args ...string) (exit bool)
}

func (c cmd) Match(s string) bool {
	return c.Name == s
}

func (c cmd) Run(w io.Writer, args ...string) bool {
	return c.Action(w, args...)
}

var cmds = make([]cmd, 0, 10)

func init() {
	cmds = append(cmds,
		cmd{
			Name:   "exit",
			Help:   "Exit the application",
			Action: exitCmd,
		},
		cmd{
			Name:   "help",
			Help:   "Shows the list of commands",
			Action: helpCmd,
		},
		cmd{
			Name:   "print",
			Help:   "Prints a file",
			Action: printCmd,
		},
	)
}

func main() {
	s := bufio.NewScanner(os.Stdin)
	w := os.Stdout
	fmt.Fprintln(w, "Welcome to my Pseudo-Terminals!")
	for {
		pwd, err := os.Getwd()
		if err != nil {
			fmt.Println("Cannot get working directory:", err)
			return
		}

		fmt.Fprintf(w, "\n[%s] > ", filepath.Base(pwd))
		if !s.Scan() {
			continue
		}

		args := strings.Split(string(s.Bytes()), " ")

		idx := -1
		for i := range cmds {
			if !cmds[i].Match(args[0]) {
				continue
			}

			idx = i
			break
		}

		if idx == -1 {
			fmt.Fprintf(w, "%q not found. Use `help` for available commands\n", args[0])
			continue
		}

		if cmds[idx].Run(w, args[1:]...) {
			fmt.Fprintln(w)
			return
		}
	}
}
