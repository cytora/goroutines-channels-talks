package main

import (
	"fmt"
	"sync"
	"time"
)


func concealerWorker(id int, items <- chan string, hat chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for item := range items {
		time.Sleep(1 * time.Second)
		fmt.Printf("Concealer %d hides %s in the hat!\n", id, item)
		time.Sleep(1 * time.Second)
		hat <- item
	}
}
func createConcealerPool(items <- chan string, hat chan<- string, noOfWorkers int) {
	var wg sync.WaitGroup // know when all workers have completed
	for i := 0; i < noOfWorkers; i++ {
		wg.Add(1)
		go concealerWorker(i, items, hat, &wg)
	}
	wg.Wait()
	fmt.Println("Concealers have finished their work!")
	close(hat) // all concealers have finished no more to conceal/reveal
}

//START OMIT
func fetchItems(itemsCh chan <- string) {
	items := []string{"bowling ball", "flowers", "rabbit", "melon", "dove", "boot"}
	for _, item := range items {
		time.Sleep(1 * time.Second)
		itemsCh <- item // populate the items channel
	}
	fmt.Println("Finished fetching items!")
	close(itemsCh) // all done with the items
}
func main() {
	const concealerCount = 3
	fmt.Println("The magic show has started!")

	itemsCh := make(chan string, concealerCount) // all concealers have an item to pick up
	hatCh := make(chan string, concealerCount) // all concealers can put down their result
	go fetchItems(itemsCh) // populate items to hide
	go createConcealerPool(itemsCh, hatCh, concealerCount) // let the concealers do their job

	for item := range hatCh {
		time.Sleep(1 * time.Second)
		fmt.Printf("The magician reveals %s from the hat!\n", item)
	}
	fmt.Println("The magic show has ended!")
}
//END OMIT
