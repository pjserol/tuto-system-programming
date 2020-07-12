package main

import (
	"fmt"
	"log"
	"sync"
)

func main() {
	log.Println("One producer one consumer")
	oneProducerOneConsumer()

	log.Println("Multple producers")
	multiplePoducers()

	log.Println("Multiple consumers")
	multipleConsumers()

	log.Println("Multiple consumers & producers")
	multipleConsumersAndProducers()
}

func oneProducerOneConsumer() {
	// 1 Producer
	var ch = make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
		close(ch)
	}()

	// 1 Consumer
	var done = make(chan struct{})
	go func() {
		for i := range ch {
			fmt.Println(i)
		}
		close(done)
	}()
	<-done
}

// multiplePoducers (N * 1)
func multiplePoducers() {
	// 3 Producer
	var ch = make(chan string)
	wg := sync.WaitGroup{}
	wg.Add(3)
	for i := 0; i < 3; i++ {
		go func(n int) {
			for i := 0; i < 5; i++ {
				ch <- fmt.Sprintln(n, i)
			}
			wg.Done()
		}(i)
	}
	go func() {
		wg.Wait()
		close(ch)
	}()

	// 1 Consumer
	var done = make(chan struct{})
	go func() {
		for i := range ch {
			fmt.Println(i)
		}
		close(done)
	}()
	<-done
}

// multipleConsumers (1 * M)
func multipleConsumers() {
	// 3 Consumers
	wg := sync.WaitGroup{}
	wg.Add(3)
	var ch = make(chan string)

	for i := 0; i < 3; i++ {
		go func(n int) {
			for i := range ch {
				fmt.Println(n, i)
			}
			wg.Done()
		}(i)
	}

	// 1 Producer
	go func() {
		for i := 0; i < 5; i++ {
			ch <- fmt.Sprintln("prod-", i)
		}
		close(ch)
	}()

	wg.Wait()
}

// multipleConsumersAndProducers (N * M)
func multipleConsumersAndProducers() {
	const (
		N = 3 // producers
		M = 5 // consumers
	)

	wg1 := sync.WaitGroup{}
	wg1.Add(N)

	wg2 := sync.WaitGroup{}
	wg2.Add(M)

	var ch = make(chan string)

	for i := 0; i < N; i++ {
		go func(n int) {
			for i := 0; i < 5; i++ {
				ch <- fmt.Sprintf("src-%d[%d]", n, i)
			}
			wg1.Done()
		}(i)
	}

	for i := 0; i < M; i++ {
		go func(n int) {
			for i := range ch {
				fmt.Printf("cons-%d, msg %q\n", n, i)
			}
			wg2.Done()
		}(i)
	}

	wg1.Wait()
	close(ch)
	wg2.Wait()
}
