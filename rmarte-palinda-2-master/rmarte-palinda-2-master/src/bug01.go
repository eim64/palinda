package main

import (
	"fmt"
)

// I want this program to print "Hello world!", but it doesn't work.
func main() {
	ch := make(chan string, 1)
	ch <- "Hello World!"
	fmt.Println(<-ch)
}
