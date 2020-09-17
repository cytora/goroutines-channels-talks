package main

import "fmt"

//START OMIT
func myFunc() {
	fmt.Println("This is the code I will run concurrently.")
}

func main() {
	// Start the goroutine using the go keyword - EASY!
	go myFunc()
}
//END OMIT