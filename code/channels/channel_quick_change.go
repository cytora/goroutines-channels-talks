package main

import (
	"fmt"
	"time"
)

//START OMIT
func quickChange(done chan bool) {
	fmt.Println("The magician wears a RED suit")
	fmt.Println("The magician's assistant draws curtain around him")
	time.Sleep(2 * time.Second)
	fmt.Println("The magician wears a BLUE suit")
	fmt.Println("The magician takes a bow")
	done <- true
}
func main() {
	fmt.Println("The magic show has started!")

	//a chan for us to know when the magician has finished performing
	done := make(chan bool)
	// our magician takes the stage
	go quickChange(done)

	// wait for the magician to finish
	<-done
	// end the magic show
	fmt.Println("The magic show has ended!")
}
//END OMIT