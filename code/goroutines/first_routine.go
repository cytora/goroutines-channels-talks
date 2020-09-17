package main

import (
	"fmt"
)

func readMinds() {
	fmt.Println("You chose the Queen of Hearts!")
}
func main() {
	// start the magic show
	fmt.Println("The magic show has started!")
	fmt.Println("Think of a playing card.")
	fmt.Println("Our magician will now read your mind!")

	// tell magician to read our minds
	go readMinds()

	// end the magic show
	fmt.Println("The magic show has ended!")
}