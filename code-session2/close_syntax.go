package main

import "fmt"

func main() {
	coins := make(chan string)
	//START OMIT
	c, ok := <- coins
	if ok {
		fmt.Printf("%s has been teleported!", c)
	}
	//END OMIT

}