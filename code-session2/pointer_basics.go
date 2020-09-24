package main

import (
	"fmt"
	"time"
)

//START OMIT
func main() {
	fmt.Println("The magic show has started!")
	time.Sleep(2 * time.Second)

	// a variable
	magic := "Abracadabra"
	// a pointer to the magic variable
	magicPtr := &magic

	fmt.Printf("The magic word is %s!\n", magic)
	time.Sleep(2 * time.Second)

	fmt.Printf("The magic pointer is %v!\n", magicPtr)
	time.Sleep(2 * time.Second)

	fmt.Printf("The dereferenced magic pointer is %s!\n",*magicPtr)
	time.Sleep(2 * time.Second)

	fmt.Println("The magic show has ended!")
}
//END OMIT
