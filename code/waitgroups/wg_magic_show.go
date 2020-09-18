package main

import (
	"fmt"
	"sync"
)
//START OMIT
func magician(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("The magician catches the bullet fired at him!")
}
func assistant(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("The assistant fires a bullet at the magician!")
}
func main() {
	fmt.Println("The magic show has started!")
	var wg sync.WaitGroup

	wg.Add(1)
	go assistant(&wg)
	wg.Wait()  // wait for the assistant to fire the gun
	wg.Add(1)
	go magician(&wg)
	wg.Wait() // wait for the magician to reveal the bullet

	// end the magic show
	fmt.Println("The magic show has ended!")
}
//END OMIT
