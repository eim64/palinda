package main

import (
	"fmt"
	"time"
)

// This program should go to 11, but it seemingly only prints 1 to 10.
func main() {
	ch := make(chan int)
	fin := make(chan bool)

	go Print(ch, fin)
	for i := 1; i <= 11; i++ {
		ch <- i
	}

	close(ch)
	select {
	case <-fin:
	}
}

// Print prints all numbers sent on the channel.
// The function returns when the channel is closed.
func Print(ch <-chan int, fin chan<- bool) {
	for n := range ch { // reads from channel until it's closed
		time.Sleep(10 * time.Millisecond) // simulate processing time
		fmt.Println(n)
	}

	fin <- true
}
