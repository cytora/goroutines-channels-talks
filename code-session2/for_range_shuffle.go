package main

import (
	"fmt"
	"time"
)

//START OMIT
var topCards = []string {"queen of hearts", "four of clubs", "jack of spades"}

func forRangeShuffle(stop chan bool, done chan bool) {
	fmt.Println("The magician starts shuffling the cards.")
	time.Sleep(1 * time.Second)

	i := 0
	for range stop{
		fmt.Printf("The magician cuts the deck.\n Topcard is %s.\n", topCards[i])
		time.Sleep(2 * time.Second)
		fmt.Println("The magician resumes shuffling the cards.")
		i++
	}

	fmt.Println("The magician stops the shuffle and puts the deck down.")
	time.Sleep(2 * time.Second)

	fmt.Println("The magician takes a bow")
	close(done) //close this channel to show we are done and unblock main
}
//END OMIT
