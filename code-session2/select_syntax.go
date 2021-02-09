package main

import "fmt"

func main() {
	cards := make(chan string)
	coins := make(chan string)
	//START OMIT
	select {
	case card := <- cards:
		fmt.Printf("%s pulled from the deck!\n", card)
	case coin:= <- coins:
		fmt.Printf("%s pulled from behind your ear!\n", coin)
	default:
		fmt.Println("No card or coin received!")
	}
	//END OMIT

}
