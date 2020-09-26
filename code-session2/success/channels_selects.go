package main

import (
	"fmt"
	"time"
)

func sawPerformer(name string, saw chan bool, done chan bool) {
	for {
		select {
		case <- saw:
			fmt.Printf("%s saws the woman.\n", name)
			time.Sleep(1 * time.Second)
			saw <- true
		case <-done:
			fmt.Printf("%s has finished sawing and takes a bow!\n", name)
			return
		}
	}
}
//START OMIT
func main() {
	fmt.Println("The magic show has started!")
	time.Sleep(2 * time.Second)
	fmt.Println("The woman gets inside the box.")
	time.Sleep(2 * time.Second)

	saw := make(chan bool, 1)
	done := make(chan bool)
	go sawPerformer("Magician", saw, done)
	go sawPerformer("Assistant", saw, done)

	//send the saw down the channel to start the sawing
	saw <- true
	// allow 10s of woman sawing
	time.Sleep(10 * time.Second)
	fmt.Println("Sawing time is over!")
	close(done)  // tell the performers to stop sawing and take a bow

	time.Sleep(2 * time.Second)
	fmt.Println("The woman wiggles her toes! Wiggle Wiggle!")
	time.Sleep(2 * time.Second)
	fmt.Println("The magic show has ended!")
}
//END OMIT
