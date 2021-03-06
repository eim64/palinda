package main

import (
	"fmt"
	"io/ioutil"
	"runtime"
	"strings"
	"time"
	"unicode"
)

const DataFile = "loremipsum.txt"

// Return the word frequencies of the text argument.
//
// Split load optimally across processor cores.
func WordCount(text string) map[string]int {
	freq := make(map[string]int)
	nsplits := runtime.NumCPU() - 1
	dch := make(chan map[string]int, nsplits)

	seglen := len(text) / nsplits

	pright := 0
	for i := 1; i < nsplits; i++ {
		index := i * seglen

		for !space_or_punct(rune(text[index])) {
			index++
		}

		go channel_wordcount(text[pright:index], dch)

		pright = index
	}

	go channel_wordcount(text[pright:], dch)

	for i := 0; i < nsplits; i++ {
		wc := <-dch

		for k, v := range wc {
			val, exist := freq[k]

			if !exist {
				freq[k] = v
				continue
			}

			freq[k] = val + v
		}
	}

	return freq
}

func channel_wordcount(text string, ch chan<- map[string]int) {
	freqs := make(map[string]int)
	for _, word := range strings.FieldsFunc(strings.ToLower(text), space_or_punct) {
		count, exist := freqs[word]
		if !exist {
			freqs[word] = 1
			continue
		}

		freqs[word] = count + 1
	}
	ch <- freqs
}

func space_or_punct(c rune) bool {
	return unicode.IsSpace(c) || unicode.IsPunct(c)
}

// Benchmark how long it takes to count word frequencies in text numRuns times.
//
// Return the total time elapsed.
func benchmark(text string, numRuns int) int64 {
	start := time.Now()
	for i := 0; i < numRuns; i++ {
		WordCount(text)
	}
	runtimeMillis := time.Since(start).Nanoseconds() / 1e6

	return runtimeMillis
}

// Print the results of a benchmark
func printResults(runtimeMillis int64, numRuns int) {
	fmt.Printf("amount of runs: %d\n", numRuns)
	fmt.Printf("total time: %d ms\n", runtimeMillis)
	average := float64(runtimeMillis) / float64(numRuns)
	fmt.Printf("average time/run: %.2f ms\n", average)
}

func main() {
	// read in DataFile as a string called data
	data, _ := ioutil.ReadFile("loremipsum.txt")

	numRuns := 100
	runtimeMillis := benchmark(string(data), numRuns)
	printResults(runtimeMillis, numRuns)
}
