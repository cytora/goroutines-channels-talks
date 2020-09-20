package main

import (
	"fmt"
	"time"
)

//START OMIT
var cardPicks = []string {"queen of hearts", "four of clubs", "jack of spades"}
func magician(cards chan string) {
	for c := range cards {
		time.Sleep(500 * time.Millisecond)
		fmt.Printf("The magician guesses you picked %s\n", c)
	}
}
func main() {
	fmt.Println("The magic show has started!")
	time.Sleep(1 * time.Second)

	cards := make(chan string) //channel for card picks to send to the magician
	go magician(cards)
	for _,p := range cardPicks {
		fmt.Printf("The volunteer picks %s\n", p)
		cards <- p // signal card pick to magician
		time.Sleep(1 * time.Second)
	}

	fmt.Println("The magic show has ended!")
}
//END OMIT