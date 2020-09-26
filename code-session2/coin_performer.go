package main

import (
	"fmt"
	"time"
)

//START OMIT
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
//END OMIT
