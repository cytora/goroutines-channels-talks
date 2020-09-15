package main

import (
	"fmt"
)

func guessCard() {
	fmt.Println("You chose the Queen of Hearts!")
}
func main() {
	go guessCard()
	fmt.Println("Think of any card and the magician will guess it.")
}