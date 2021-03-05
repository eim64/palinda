package main

import (
	"fmt"
	"math/bits"
)

// sum the numbers in a and send the result on res.
func sum(a []int, res chan<- int) {
	sum := 0
	for _, val := range a {
		sum += val
	}

	res <- sum
}

// concurrently sum the array a.
func ConcurrentSum(a []int) int {
	n := len(a)
	ch := make(chan int)
	splits := 32 - bits.LeadingZeros32(uint32(n)) //budget log2
	flen := n / splits

	for s := 0; s < splits-1; s++ {
		go sum(a[s*flen:(s+1)*flen], ch)
	}

	go sum(a[(splits-1)*flen:], ch)

	// TODO Get the subtotals from the channel and return their sum
	tot := 0
	for s := 0; s < splits; s++ {
		tot += <-ch
	}

	return tot
}

func main() {
	// example call
	a := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(ConcurrentSum(a))
}
