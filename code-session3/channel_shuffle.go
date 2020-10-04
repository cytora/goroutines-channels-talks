package main

import (
	"fmt"
	"sync"
	"time"
)

//START OMIT
var tops = []string {"queen of hearts", "four of clubs", "jack of spades", "ace of diamonds"}
func channelShuffle(cardPicks chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("The magician starts shuffling the cards.")
	time.Sleep(1 * time.Second)

	for i := 0; i < len(tops); i++{
		time.Sleep(2 * time.Second)
		cardPicks <- tops[i] //send the topcard back
	}

	fmt.Println("The magician takes a bow!")
}
//END OMIT