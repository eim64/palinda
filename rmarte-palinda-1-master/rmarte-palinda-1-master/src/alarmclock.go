package main

import (
	"fmt"
	"time"
)

func Reminder(interval int, message string) {
	timer := time.NewTicker(time.Second * time.Duration(interval))

	for {
		select {
		case t := <-timer.C:
			fmt.Println("The time is now ", t.Format("15.04.05"), ": ", message)
		}
	}
}

func main() {
	go Reminder(10, "Time to eat")
	go Reminder(30, "Time to work")
	go Reminder(60, "Time to sleep")

	select {}
}
