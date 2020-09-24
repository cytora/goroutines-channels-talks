package main

import (
	"fmt"
	"time"
)

//START OMIT
var topCards = []string {"queen of hearts", "four of clubs", "jack of spades"}
func shuffle(signals chan bool, done chan bool) {
	var pcount int
	for range signals {
		fmt.Printf("The magician shuffles the deck.\n Topcard is %s.\n", topCards[pcount])
		pcount++
	}
	fmt.Println("The magician takes a bow")
	close(done) //close this channel to show we are done and unblock main
}
func main() {
	fmt.Println("The magic show has started!")
	signals := make(chan bool) //channel for card picks to send to the magician
	done := make(chan bool) //channel to signal that the magician finished bowing
	go shuffle(signals, done)
	for range topCards {
		time.Sleep(2 * time.Second)
		signals <- true //signal to stop the shuffle and reveal the top card
	}
	close(signals)
	<-done
	fmt.Println("The magic show has ended!")
}
//END OMIT