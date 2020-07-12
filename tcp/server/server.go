package main

import (
	"bufio"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"time"
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

	listener, err := net.ListenTCP("tcp", addr)
	if err != nil {
		log.Fatalln("Listener:", address, err)
	}

	for {
		time.Sleep(time.Millisecond * 100)

		conn, err := listener.AcceptTCP()
		if err != nil {
			log.Fatalln("<- Accept:", address, err)
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	r := bufio.NewReader(conn)
	time.Sleep(time.Second / 2)
	for {
		msg, err := r.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				log.Println("<-", err)
				return
			}
			if nerr, ok := err.(net.Error); ok && !nerr.Temporary() {
				log.Println("<- Network error:", err)
				return
			}
			log.Println("<- Message error:", err)
			continue
		}
		switch msg = strings.TrimSpace(msg); msg {
		case `\q`:
			log.Println("Exit!")
			if err := conn.Close(); err != nil {
				log.Println("<- Close:", err)
			}
			time.Sleep(time.Second / 2)
			return
		case `\s`:
			log.Println("<- This a secret message!")
		default:
			log.Println("<- Message Received:", msg)
		}
	}
}
