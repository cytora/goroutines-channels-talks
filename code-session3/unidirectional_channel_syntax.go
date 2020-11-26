package main

import (
	"fmt"
	"time"
)

//START OMIT
func receive(r <-chan string) {
	fmt.Printf("The type of the receive channel is %T\n", r)
}
func send(r chan<- string) {
	fmt.Printf("The type of the send channel is %T\n", r)
}
func main() {
	fmt.Println("The magic show has started!")
	time.Sleep(2 * time.Second)

	ch := make(chan string)
	fmt.Printf("The type of the initial channel is %T\n", ch)
	time.Sleep(2 * time.Second)
	receive(ch)
	time.Sleep(2 * time.Second)
	send(ch)

	fmt.Println("The magic show has ended!")
}
//END OMIT
