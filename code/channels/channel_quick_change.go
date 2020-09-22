package main

import (
	"fmt"
	"time"
)

//START OMIT
func quickChange(done chan bool) {
	fmt.Println("The magician wears a RED suit!")
	time.Sleep(500 * time.Millisecond)

	fmt.Println("The magician's assistant draws curtain around him")
	time.Sleep(3 * time.Second)

	fmt.Println("The magician wears a BLUE suit!")
	time.Sleep(500 * time.Millisecond)

	fmt.Println("The magician takes a bow")
	done <- true //the magician writes to the done channel
}
func main() {
	fmt.Println("The magic show has started!")
	time.Sleep(1 * time.Second)

	//a chan for us to know when the magician has finished performing
	done := make(chan bool)
	go quickChange(done)
	<-done

	fmt.Println("The magic show has ended!")
}
//END OMIT