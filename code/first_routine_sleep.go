package main

import (
	"fmt"
	"time"
)

func readMindsAgain() {
	fmt.Println("You chose the Queen of Hearts!")
}
func main() {
	// start the magic show
	fmt.Println("Think of any card and the magician will read your mind.")
	// tell magician to read our minds
	go readMindsAgain()
	// give the magician some time to think
	time.Sleep(5 * time.Second)
	// end the magic show
	fmt.Println("The magic show has ended!")
}