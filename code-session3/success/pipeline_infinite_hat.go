package main

import (
	"fmt"
	"sync"
	"time"
)

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
	fmt.Printf("The concealer %d takes a bow!\n", id)
}
//START OMIT
const revealerCount int = 2
var items []string = []string{"bowling ball", "flowers", "rabbit", "melon"}
func main() {
	fmt.Println("The magic show has started!")
	time.Sleep(2 * time.Second)

	var concealersWg, revealersWg sync.WaitGroup //waitgroups for the  concealers & revealers
	contents := make(chan string, revealerCount) //buffer items to equal to revealer count
	revealersWg.Add(revealerCount)
	for i := 0; i< revealerCount; i++ {
		go revealer(i, contents, &revealersWg)
	}
	concealersWg.Add(len(items))
	for i, item := range items {
		go concealer(i, item, contents, &concealersWg)
		time.Sleep(1 * time.Second)
	}
	concealersWg.Wait() //wait for the concealers to finis
	close(contents) //signal to revealers no more contents are coming
	revealersWg.Wait()//wait for the revealers for finish emptying the hat

	time.Sleep(2 * time.Second)
	fmt.Println("The magic show has ended!")
}
//END OMIT
