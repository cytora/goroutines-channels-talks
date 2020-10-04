package main

import (
	"fmt"
	"sync"
	"time"
)

var tops = []string {"queen of hearts", "four of clubs", "jack of spades", "ace of diamonds"}
func channelShuffle(cardPicks chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("The magician starts shuffling the cards.")
	time.Sleep(1 * time.Second)

	for i := 0; i < len(tops); i++{
		time.Sleep(2 * time.Second)
		cardPicks <- tops[i] //send the topcard back
	}

	fmt.Println("The magician takes a bow!")
}
//START OMIT
func main() {
	fmt.Println("The magic show has started!")
	time.Sleep(2 * time.Second)

	var wg sync.WaitGroup
	cards := make(chan string) //channel to receive card picks on
	wg.Add(1)
	go channelShuffle(cards, &wg)

	// receive 2 cards from from the magician and call it a day
	fmt.Printf("The magician reveals the %s!\n", <- cards)
	fmt.Printf("The magician reveals the %s!\n", <- cards)

	time.Sleep(1 * time.Second)
	fmt.Println("No more cards, thanks!")
	close(cards) // we don't want to receive top cards anymore
	wg.Wait()    // wait for magician to finish his work

	fmt.Println("The magic show has ended!")
}
//END OMIT