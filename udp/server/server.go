package main

import (
	"bytes"
	"log"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("Please specify an address!")
	}

	address := os.Args[1]

	addr, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		log.Fatalln("Invalid address:", address, err)
	}
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Fatalln("Listener:", address, err)
	}

	b := make([]byte, 1024)
	for {
		n, addr, err := conn.ReadFromUDP(b)
		if err != nil {
			log.Println("<-", addr, "Message error:", err)
			continue
		}
		msg := bytes.TrimSpace(b[:n])
		log.Printf("<- %q from %s", msg, addr)
		for i, l := 0, len(msg); i < l/2; i++ {
			msg[i], msg[l-1-i] = msg[l-1-i], msg[i]
		}

		msg = append(msg, '\n')
		if _, err := conn.WriteTo(b[:n], addr); err != nil {
			log.Println("->", addr, "Send error:", err)
		}
	}
}
