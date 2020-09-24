package main

import (
	"fmt"
	"sync"
	"time"
)

//START OMIT
func quickChange(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("The magician wears a RED suit!")
	time.Sleep(500 * time.Millisecond)

	fmt.Println("The assistant draws the curtain around him")
	time.Sleep(3 * time.Second)

	fmt.Println("The magician wears a BLUE suit!")
	time.Sleep(500 * time.Millisecond)

	fmt.Println("The magician takes a bow")
}
func main() {
	fmt.Println("The magic show has started!")
	time.Sleep(1 * time.Second)

	var wg sync.WaitGroup // wait for the magician to complete
	wg.Add(1)
	go quickChange(&wg)
	wg.Wait()

	fmt.Println("The magic show has ended!")
}
//END OMIT