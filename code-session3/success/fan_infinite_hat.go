package main

import (
	"fmt"
	"sync"
	"time"
)

func concealer(id int, items <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		for item := range items {
			time.Sleep(1 * time.Second)
			fmt.Printf("The concealer %d puts %s in the hat!\n", id, item)
			out <- item
			time.Sleep(1 * time.Second)
		}
		fmt.Printf("The concealer %d takes a bow.\n", id)
		close(out)
	}()
	return out
}
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
		close(out) //close the channel once all items channels are drained
	}()
	return out
}
//START OMIT
func main() {
	fmt.Println("The magic show has started!")
	items := []string{"bowling ball", "flowers", "rabbit", "melon"}
	in := make(chan string)
	go func() { // convert items list to channel
		for _, item := range items{
			time.Sleep(1 * time.Second)
			in <- item
		}
		close(in) // close channel when items finished
	}()

	// Two concealers will put items in the hat - FAN OUT
	c1 := concealer(0, in)
	c2 := concealer(1, in)
	// One revealer will pull the items out of all the channels - FAN IN
	for c:= range revealer(c1,c2) {
		time.Sleep(1 * time.Second)
		fmt.Printf("The revealer pulls %s out of the hat!\n", c)
	}
	fmt.Println("The revealer takes his bow!")
	fmt.Println("The magic show has ended!")
}
//END OMIT
