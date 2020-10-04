package main

import (
	"fmt"
	"sync"
	"time"
)

//START OMIT
func revealer(id int, contents <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	//the revealer receives items from the other performers and pulls them out of the hat
	for c:= range contents {
		time.Sleep(1 * time.Second)
		fmt.Printf("The revealer %d pulls %s out of the hat!\n", id, c)
	}

	time.Sleep(1 * time.Second)
	fmt.Printf("The revealer %d takes a bow!\n", id)
}
func concealer(id int, item string, contents chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()

	//the concealers put contents in the hat
	time.Sleep(1 * time.Second)
	fmt.Printf("The concealer %d puts %s in the hat!\n", id, item)
	contents <- item
	time.Sleep(1 * time.Second)
}
  //END OMIT
