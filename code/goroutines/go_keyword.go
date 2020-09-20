package main

import "fmt"

//START OMIT
func myFunc() {
	fmt.Println("This is the code I will run concurrently.")
}
func main() {
	go myFunc() //start the goroutine using the go keyword - EASY!
}
//END OMIT