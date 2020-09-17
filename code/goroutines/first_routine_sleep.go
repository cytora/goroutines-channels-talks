package main

import (
	"fmt"
	"time"
)
func readMindsAgain() {
	fmt.Println("You chose the Queen of Hearts!")
}

//START OMIT
func main() {
	// start the magic show
	fmt.Println("The magic show has started!")
	fmt.Println("Think of a playing card.")
	fmt.Println("Our magician will now read your mind!")

	// tell magician to read our minds
	go readMindsAgain()

	// give the magician some time to think
	time.Sleep(5 * time.Second)

	// end the magic show
	fmt.Println("The magic show has ended!")
}
//END OMIT