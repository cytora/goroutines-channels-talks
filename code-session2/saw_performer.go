package main

import (
	"fmt"
	"time"
)

//START OMIT
func sawPerformer(name string, saw chan bool, done chan bool) {
	for {
		select {
		case <- saw:
			fmt.Printf("%s saws the woman.\n", name)
			time.Sleep(1 * time.Second)
			saw <- true
		case <-done:
			time.Sleep(1 * time.Second)
			fmt.Printf("%s has finished sawing and takes a bow!\n", name)
			return
		}
	}
}
//END OMIT