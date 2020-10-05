package main

import (
	"fmt"
	"time"
)

//START OMIT
func assistant(actions []string) <-chan int {
	out := make(chan int)
	go func() {
		for i, item := range actions {
			fmt.Printf("The assistant %s!\n", item)
			time.Sleep(2 * time.Second)
			out <- i
		}
		close(out)
	}()
	return out
}
func magician(items <-chan int) <-chan string {
	magicianActions := []string{"unlocks jacket", "unties jacket", "takes off jacket"}
	out := make(chan string)
	go func() {
		for i := range items {
			time.Sleep(2 * time.Second)
			out <- magicianActions[i]
		}
		close(out)
	}()
	return out
}
//END OMIT