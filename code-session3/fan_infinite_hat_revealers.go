package main

import (
	"sync"
	"time"
)

//START OMIT
func revealer(cs ...<-chan string) <-chan string {
	var wg sync.WaitGroup
	out := make(chan string)
	//Start one goroutine per channel to merge all the items into 1
	wg.Add(len(cs))
	for _, c := range cs {
		go func(c <-chan string) {
			defer wg.Done()
			for n := range c { //drain each items channel
				time.Sleep(1 * time.Second)
				out <- n
			}
		}(c)
	}
	go func() {
		wg.Wait()
		time.Sleep(1 * time.Second)
		close(out) //close the channel once all items channels are drained
	}()
	return out
}
//END OMIT