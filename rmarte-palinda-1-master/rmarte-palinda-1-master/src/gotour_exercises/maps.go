package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	ret := map[string]int{}
	for _, w := range strings.Split(s, " ") {
		val, exists := ret[w]

		if !exists {
			ret[w] = 1
			continue
		}

		ret[w] = val + 1
	}

	return ret
}

func main() {
	wc.Test(WordCount)
}
