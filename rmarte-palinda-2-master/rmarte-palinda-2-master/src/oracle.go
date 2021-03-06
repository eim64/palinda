// Stefan Nilsson 2013-03-13

// This program implements an ELIZA-like oracle (en.wikipedia.org/wiki/ELIZA).
package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

const (
	star   = "Pythia"
	venue  = "Delphi"
	prompt = "> "
)

func main() {
	fmt.Printf("Welcome to %s, the oracle at %s.\n", star, venue)
	fmt.Println("Your questions will be answered in due time.")

	questions := Oracle()
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(prompt)
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		fmt.Printf("%s heard: %s\n", star, line)
		questions <- line // The channel doesn't block.
	}
}

// Oracle returns a channel on which you can send your questions to the oracle.
// You may send as many questions as you like on this channel, it never blocks.
// The answers arrive on stdout, but only when the oracle so decides.
// The oracle also prints sporadic prophecies to stdout even without being asked.
func Oracle() chan<- string {
	questions := make(chan string)
	results := make(chan string)

	go answer(questions, results)
	go jargon(results)
	go print(results)

	return questions
}

func jargon(results chan<- string) {
	for {
		time.Sleep(time.Duration(20+rand.Intn(40)) * time.Second)
		results <- "My spirtualitymeter 3000 is showing a tremendous reaction!!!"
		time.Sleep(time.Duration(2) * time.Second)
		results <- "There's a ghost nearby, it says the string of its name satisfies [A - z], is it a relative perhaps?"
		time.Sleep(time.Duration(5) * time.Second)
		results <- "From its vibrations I can deduce that it's either male of female, probably a very close relative as well."
		time.Sleep(time.Duration(3) * time.Second)
		results <- "Using it as a medium to the great unknown, I can foresee that you will leave as an unsatisfied customer"
	}
}

func answer(questions <-chan string, results chan<- string) {
	for question := range questions {
		prophecy(question, results)
	}
}

func print(results <-chan string) {
	for res := range results {
		fmt.Println(res)
	}
}

// This is the oracle's secret algorithm.
// It waits for a while and then sends a message on the answer channel.
// TODO: make it better.
func prophecy(question string, answer chan<- string) {
	// Keep them waiting. Pythia, the original oracle at Delphi,
	// only gave prophecies on the seventh day of each month.
	time.Sleep(time.Duration(2+rand.Intn(3)) * time.Second)

	// Find the longest word.
	longestWord := ""
	words := strings.Fields(question) // Fields extracts the words into a slice.
	for _, w := range words {
		if len(w) > len(longestWord) {
			longestWord = w
		}
	}

	// Cook up some pointless nonsense.
	nonsense := []string{
		"The moon is dark.",
		"The sun is bright.",
	}
	answer <- longestWord + "... " + nonsense[rand.Intn(len(nonsense))]
}

func init() { // Functions called "init" are executed before the main function.
	// Use new pseudo random numbers every time.
	rand.Seed(time.Now().Unix())
}
