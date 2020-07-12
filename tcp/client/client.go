package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("Please specify an address!")
	}

	address := os.Args[1]
	addr, err := net.ResolveTCPAddr("tcp", address)
	if err != nil {
		log.Fatalln("Invalid address:", address, err)
	}

	createConn(addr)
}

func createConn(addr *net.TCPAddr) {
	conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		log.Fatalln("-> Connection:", err)
	}

	log.Println("-> Connection to", addr)

	r := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("# ")
		msg, err := r.ReadBytes('\n')
		if err != nil {
			log.Println("-> Message error:", err)
		}
		if _, err := conn.Write(msg); err != nil {
			log.Println("-> Connection:", err)
			return
		}
	}
}
