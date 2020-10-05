package main

import (
	"fmt"
	"time"
)

//START OMIT
func concealer(id int, items <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		for item := range items {
			time.Sleep(1 * time.Second)
			fmt.Printf("The concealer %d puts %s in the hat!\n", id, item)
			time.Sleep(1 * time.Second)
			out <- item
		}
		fmt.Printf("The concealer %d takes a bow.\n", id)
		close(out)
	}()
	return out
}
//END OMIT