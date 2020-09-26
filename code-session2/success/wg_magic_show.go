package main

import (
	"fmt"
	"sync"
	"time"
)

//START OMIT
func magician(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("The magician catches the bullet fired at him!")
	time.Sleep(500 * time.Millisecond)
}
func assistant(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(500 * time.Millisecond)
	fmt.Println("The assistant fires a bullet at the magician!")
}
func main() {
	fmt.Println("The magic show has started!")
	time.Sleep(1 * time.Second)
	var wg sync.WaitGroup

	wg.Add(1)
	go assistant(&wg) //remember to pass a pointer to the WaitGroup here
	wg.Wait()  // wait for the assistant to fire the gun
	wg.Add(1)
	go magician(&wg) //remember to pass a pointer to the WaitGroup here
	wg.Wait() // wait for the magician to reveal the bullet

	fmt.Println("The magic show has ended!")
}
//END OMIT
