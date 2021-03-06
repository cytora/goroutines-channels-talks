package main

import (
	"fmt"
	"sync"
	"time"
)

func magician(wg sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Magician's WaitGroup is %p\n,", &wg)
	fmt.Println("The magician catches the bullet fired at him!")
	time.Sleep(500 * time.Millisecond)
}

func assistant(wg sync.WaitGroup) {
	defer wg.Done()

	time.Sleep(500 * time.Millisecond)
	fmt.Printf("Assistant's WaitGroup is %p\n", &wg)
	fmt.Println("The assistant fires a bullet at the magician!")
}
//START OMIT
func main() {
	fmt.Println("The magic show has started!")
	time.Sleep(1 * time.Second)

	var wg sync.WaitGroup
	fmt.Printf("Main's WaitGroup is %p\n", &wg)
	wg.Add(1)
	go assistant(wg)
	wg.Wait()  // wait for the assistant to fire the gun

	wg.Add(1)
	go magician(wg)
	wg.Wait() // wait for the magician to reveal the bullet

	fmt.Println("The magic show has ended!")
}
//END OMIT
