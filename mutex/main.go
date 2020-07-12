package main

import (
	"fmt"
	"sync"
)

func main() {
	notThreadSafe()
	threadSafe()
	valueEmbedded()
}

func notThreadSafe() {
	done := make(chan struct{}, 10000)
	var a = 0

	for i := 0; i < cap(done); i++ {
		go func(i int) {
			if i%2 == 0 {
				a++
			} else {
				a--
			}

			done <- struct{}{}
		}(i)
	}

	for i := 0; i < cap(done); i++ {
		<-done
	}

	fmt.Println(a)
}

func threadSafe() {
	done := make(chan struct{}, 10000)
	m := sync.Mutex{}
	var a = 0

	for i := 0; i < cap(done); i++ {
		go func(i int, l sync.Locker) {
			l.Lock()
			defer l.Unlock()
			if i%2 == 0 {
				a++
			} else {
				a--
			}

			done <- struct{}{}
		}(i, &m)
	}

	for i := 0; i < cap(done); i++ {
		<-done
	}

	fmt.Println(a)
}

type counter struct {
	m     sync.Mutex
	value int
}

func (c *counter) Increment() {
	c.m.Lock()
	c.value++
	c.m.Unlock()
}

func (c *counter) Decrement() {
	c.m.Lock()
	c.value--
	c.m.Unlock()
}

func (c *counter) Value() int {
	c.m.Lock()
	a := c.value
	c.m.Unlock()
	return a
}

func valueEmbedded() {
	done := make(chan struct{}, 10000)
	var a = counter{}

	for i := 0; i < cap(done); i++ {
		go func(i int) {
			if i%2 == 0 {
				a.Increment()
			} else {
				a.Decrement()
			}
			done <- struct{}{}
		}(i)
	}

	for i := 0; i < cap(done); i++ {
		<-done
	}

	fmt.Println(a.Value())
}
