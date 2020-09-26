package main

import (
	"fmt"
	"time"
)

//START OMIT
func sawPerformer(name string, saw chan bool, done chan bool) {
	for {
		select {
		case <- saw:
			fmt.Printf("%s saws the woman.\n", name)
			time.Sleep(1 * time.Second)
			saw <- true
		case <-done:
			fmt.Printf("%s has finished sawing and takes a bow!\n", name)
			time.Sleep(1 * time.Second)
		}
	}
}
func main() {
	fmt.Println("The magic show has started!")
	time.Sleep(2 * time.Second)
	fmt.Println("The woman gets inside the box.")

	saw := make(chan bool)
	done := make(chan bool)
	go sawPerformer("Magician", saw, done)
	go sawPerformer("Assistant", saw, done)

	// allow 10s of woman sawing
	time.Sleep(10 * time.Second)
	done <- true

	fmt.Println("The woman wiggles her toes! Wiggle Wiggle!")
	time.Sleep(2 * time.Second)
	fmt.Println("The magic show has ended!")
}
//END OMIT
