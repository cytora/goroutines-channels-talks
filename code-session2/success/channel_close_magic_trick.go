package main

import (
	"fmt"
	"time"
)

var tops = []string {"queen of hearts", "four of clubs", "jack of spades"}

func channelShuffle(stop chan bool, done chan bool) {
	fmt.Println("The magician starts shuffling the cards.")
	time.Sleep(1 * time.Second)

	i := 0
	for {
		_, ok := <- stop
		if !ok {
			fmt.Println("The magician stops the shuffle and puts the deck down.")
			break
		}

		fmt.Printf("The magician cuts the deck.\n Topcard is %s.\n", tops[i])
		time.Sleep(2 * time.Second)
		fmt.Println("The magician resumes shuffling the cards.")
		i++
	}

	fmt.Println("The magician takes a bow")
	close(done) //close this channel to show we are done and unblock main
}

//START OMIT
func main() {
	fmt.Println("The magic show has started!\n For this trick we will need a volunteer.")
	time.Sleep(2 * time.Second)

	stop := make(chan bool) //channel for card picks to send to the magician
	done := make(chan bool) //channel to signal that the magician finished bowing
	go channelShuffle(stop, done)

	for i := 0; i< len(tops); i++ {
		time.Sleep(2 * time.Second)
		fmt.Println("The volunteer says to cut the deck.")
		stop <- true //signal to stop the shuffle and reveal the top card
	}

	close(stop) // the volunteer is done picking cards
	<-done // the magician is done too

	time.Sleep(2 * time.Second)
	fmt.Println("The magic show has ended!")
}
//END OMIT
