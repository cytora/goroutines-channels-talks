package main

import (
	"fmt"
	"time"
)

var magicianActions = []string{"unlocks jacket", "unties jacket", "takes off jacket"}
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
//START OMIT
var assistantActions = []string{"puts jacket on magician", "ties jacket", "locks jacket"}
func main() {
	fmt.Println("The magic show has started!")
	time.Sleep(2 * time.Second)

	actions := assistant(assistantActions)
	out := magician(actions)

	for item := range out {
		time.Sleep(1 * time.Second)
		fmt.Printf("The magician %s\n", item)
		time.Sleep(1 * time.Second)
	}

	time.Sleep(2 * time.Second)
	fmt.Println("The magic show has ended!")
}
//END OMIT
