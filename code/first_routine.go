package main

import (
	"fmt"
)

func readMinds() {
	fmt.Println("You chose the Queen of Hearts!")
}
func main() {
	// start the magic show
	fmt.Println("Think of any card and the magician will read your mind.")
	// tell magician to read our minds
	go readMinds()
	// end the magic show
	fmt.Println("The magic show has ended!")
}