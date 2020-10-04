package main

import (
	"fmt"
	"sync"
	"time"
)

var topCards = []string {"queen of hearts", "four of clubs", "jack of spades", "ace of diamonds"}
func channelSignalShuffle(cardPicks chan<- string, signal chan struct{}, wg *sync.WaitGroup) {
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
//START OMIT
func main() {
	fmt.Println("The magic show has started!")
	time.Sleep(2 * time.Second)

	var wg sync.WaitGroup
	cards := make(chan string) //channel to receive card picks on
	signal := make(chan struct{}) //channel to signal shuffle stop
	wg.Add(1)
	go channelSignalShuffle(cards, signal, &wg)

	// receive 2 cards from from the magician and call it a day
	fmt.Printf("The magician reveals the %s!\n", <- cards)
	fmt.Printf("The magician reveals the %s!\n", <- cards)

	time.Sleep(1 * time.Second)
	fmt.Println("No more cards, thanks!")
	close(signal) // close signal to tell the magician to stop the shufflw
	wg.Wait()    // wait for magician to finish his work

	fmt.Println("The magic show has ended!")
}
//END OMIT