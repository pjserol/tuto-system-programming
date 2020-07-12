package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.NewTicker(time.Millisecond)
	stop := time.NewTimer(time.Millisecond * 3)

	for {
		select {
		case a := <-tick.C:
			fmt.Println(a)
		case <-stop.C:
			tick.Stop()
		case <-time.After(time.Millisecond):
			return
		}
	}
}
