package main

import (
	"fmt"
	"sync"
	"time"
)
type InfiniteHat struct {
	contents string
}
func (h *InfiniteHat) placeItem(s string) {
	h.contents = fmt.Sprintf("%s\n%s", h.contents, s)
}

//START OMIT
var items = []string{"rabbit", "bottle", "handkerchief", "bowling ball", "flowers"}
func hatMagician(item string, wg *sync.WaitGroup, hat *InfiniteHat) {
	defer wg.Done()
	hat.placeItem(item)
	fmt.Printf("The perfomer places %s in the hat\n", item)
	time.Sleep(500 * time.Millisecond)
}
func main() {
	fmt.Println("The magic show has started!")
	time.Sleep(2 * time.Second)

	hat := &InfiniteHat{}
	var wg sync.WaitGroup
	wg.Add(len(items))
	for _, item := range items {
		go hatMagician(item, &wg, hat)
	}
	wg.Wait() //wait for the performers to finish their tricks

	fmt.Printf("Hat contents after the trick are:\n%s\n", hat.contents)
	time.Sleep(1 * time.Second)
	fmt.Println("The magic show has ended!")
}
//END OMIT