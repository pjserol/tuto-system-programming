package main

import (
	"fmt"
	"time"
)

type sem chan struct{}

func (s sem) Acquire() {
	s <- struct{}{}
}

func (s sem) Release() {
	<-s
}

func main() {
	s := make(sem, 3)
	for i := 0; i < 5; i++ {
		go func(i int) {
			s.Acquire()
			fmt.Println(i, "start")
			time.Sleep(time.Second)
			fmt.Println(i, "end")
			s.Release()
		}(i)
	}
	time.Sleep(time.Second * 2)
}
