package main

import (
	"fmt"
	"time"
)
//START OMIT
func magician() {
	fmt.Println("The magician catches the bullet fired at him!")
}
func assistant() {
	fmt.Println("The assistant fires a bullet at the magician!")
}
func main() {
	fmt.Println("The magic show has started!")
	time.Sleep(2 * time.Second)

	go assistant()
	go magician()

	// same as before - give our performers time to perform
	time.Sleep(5 * time.Second)
	// close down the show
	fmt.Println("The magic show has ended!")
}
//END OMIT