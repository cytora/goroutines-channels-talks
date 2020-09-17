package main

import (
	"fmt"
	"sync"
)
//START OMIT
func trapeze(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("The acrobat jumps the trapeze.")
}
func contortionist(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("The contortionist takes the stage.")
}
func main() {
	fmt.Println("The magic show has started!")
	var wg sync.WaitGroup

	// tell our performers it's time for the show
	wg.Add(3)
	go trapeze(&wg)
	go trapeze(&wg)
	go contortionist(&wg)
	wg.Wait() // wait for our performers to finish

	// end the magic show
	fmt.Println("The magic show has ended!")
}
//END OMIT
