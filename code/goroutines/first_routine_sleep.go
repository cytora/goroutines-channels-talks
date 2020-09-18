package main

import (
	"fmt"
	"time"
)
//START OMIT
func readMindsAgain() {
	time.Sleep(500 * time.Millisecond)
	fmt.Println("You chose the Queen of Hearts!")
}
func main() {
	fmt.Println("The magic show has started!")
	time.Sleep(1 * time.Second)
	fmt.Println("Think of a playing card.")
	time.Sleep(1 * time.Second)
	fmt.Println("Our magician will now read your mind!")

	// tell magician to read our minds
	go readMindsAgain()

	// give the magician some time to think
	time.Sleep(5 * time.Second)

	// end the magic show
	fmt.Println("The magic show has ended!")
}
//END OMIT