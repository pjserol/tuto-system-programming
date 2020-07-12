package main

import (
	"fmt"
)

func send(ch chan<- int, max int) {
	for i := 0; i < max; i++ {
		ch <- i
	}
	close(ch)
}

func receive(ch <-chan int) {
	for v := range ch {
		fmt.Println(v)
	}
}

func main() {
	var a = make(chan int)

	go send(a, 10)

	receive(a)
}
