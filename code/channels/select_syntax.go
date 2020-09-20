package main

import "fmt"

func main() {
	chan1 := make(chan string)
	chan2 := make(chan string)
	//START OMIT
	select {
	case in1 := <-chan1:
		fmt.Println(in1)
	case in2 := <-chan2:
		fmt.Println(in2)
	default:
		fmt.Println("no value received")
	}
	//END OMIT

}
