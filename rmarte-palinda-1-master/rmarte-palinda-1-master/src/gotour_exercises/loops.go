package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	low := 1.0
	high := x
	var mid float64

	for i := 0; i < 10; i++ {
		mid = (low + high) / 2
		if mid*mid > x {
			high = mid
			continue
		}

		low = mid
	}

	for i := 0; i < 10; i++ {
		mid -= (mid*mid - x) / (2 * mid)
	}

	return mid
}

func main() {
	fmt.Println(Sqrt(2))
}
