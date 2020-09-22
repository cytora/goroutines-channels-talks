package main

import (
	"fmt"
	"sync"
	"time"
)

type MutexInfiniteHat struct {
	m *sync.Mutex
	contents string
}

func (h *MutexInfiniteHat) placeItem(s string) {
	h.m.Lock() //acquire control of the hat
	h.contents = fmt.Sprintf("%s\n%s", h.contents, s)
	h.m.Unlock() //relinquish control of the hat once the item is placed
}

//START OMIT
var items = []string{"rabbit", "bottle", "handkerchief", "bowling ball", "flowers"}
func hatMagician(item string, wg *sync.WaitGroup, hat *MutexInfiniteHat) {
	defer wg.Done()
	hat.placeItem(item)
	fmt.Printf("The perfomer places %s in the hat\n", item)
	time.Sleep(500 * time.Millisecond)
}
func main() {
	fmt.Println("The magic show has started!")
	time.Sleep(2 * time.Second)

	var wg sync.WaitGroup
	var m sync.Mutex //just like the WaitGroup the mutex need not be initialized
	hat := &MutexInfiniteHat{m: &m} //initialize the hat using a mutex
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
