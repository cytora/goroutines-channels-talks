package main

import (
	"fmt"
	"sync"
	"time"
)

//START OMIT
var topCards = []string {"queen of hearts", "four of clubs", "jack of spades", "ace of diamonds"}
func channelSignalShuffle(cardPicks chan<- string, signal <-chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("The magician starts shuffling the cards.")
	time.Sleep(1 * time.Second)

	i := 0
	for {
		select {
		case cardPicks <- topCards[i]:
			i++
			time.Sleep(2 * time.Second) // do nothing other than sleep here
		case <- signal:
			close(cardPicks) // we won't be sending any more cards back
			time.Sleep(2 * time.Second)
			fmt.Println("The magician takes a bow!")
			return

		}
	}
}
//END OMIT