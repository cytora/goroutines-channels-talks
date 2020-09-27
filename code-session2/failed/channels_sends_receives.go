package main

import (
	"fmt"
	"time"
)

type CoinPerformer struct {
	name string
	passes int
}
func (c *CoinPerformer) perform(coin chan bool, done chan string) {
	fmt.Printf("%s ready to start the teleportation trick.\n", c.name)

	for i := 0; i < c.passes; i ++ {

		// receive coin from other performer, blocking
		<- coin

		fmt.Printf("%s has the coin!\n", c.name)
		time.Sleep(2 * time.Second)

		// send the coin to the other performer, blocking
		coin <- true
	}

	//the performer has finished and is ready to tell main they are done, blocking
	done <- c.name
}

//START OMIT
func main() {
	fmt.Println("The magic show has started!")
	coin := make(chan bool) //create the coin synchronization channel
	done := make(chan string) //create the done channel

	time.Sleep(2 * time.Second)
	fmt.Println("The teleportation trick is ready to begin!")
	magician := CoinPerformer{name: "Magician", passes: 3}
	assistant := CoinPerformer{name: "Assistant", passes: 3}
	go magician.perform(coin, done)
	go assistant.perform(coin, done)
	coin <- true

	// wait for the 2 performers to finish
	for i := 0; i < 2; i++ {
		// name of performer who has finished, blocking
		name := <- done
		fmt.Printf("%s has finished and takes a bow!\n", name)
		time.Sleep(2 * time.Second)
	}
	fmt.Println("The magic show has ended!")
}
//END OMIT

