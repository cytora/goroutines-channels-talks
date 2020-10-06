package main

import (
	"fmt"
	"sync"
	"time"
)

//START OMIT
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
//END OMIT
