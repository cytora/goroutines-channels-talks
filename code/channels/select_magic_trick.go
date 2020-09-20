package main

import (
	"fmt"
	"sync"
	"time"
)

//START OMIT
var cupPicks = []string {"cup 1", "cup 2", "cup 1", "cup 3", "cup 2", "cup 3"}
func cupsAndBalls(cupPicks <-chan string, done <-chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case c := <-cupPicks:
			fmt.Printf("The magician reveals %s empty and mixes cups!\n", c)
		case <-done:
			fmt.Printf("The magician takes a bow!\n")
			return
		}}}
func main() {
	var wg sync.WaitGroup // WaitGroup to know when the magician has bowed
	cups := make(chan string) //channel for card picks to send to the magician
	done := make(chan bool) //channel for card picks to send to the magician
	wg.Add(1)
	go cupsAndBalls(cups, done, &wg)
	for _,pick := range cupPicks {
		fmt.Printf("The volunteer picks %s\n", pick)
		cups <- pick //point out cup to the magician
		time.Sleep(1 * time.Second)
	}
	done<-true //signal to the magician that the volunteer has given up
	wg.Wait()}
//END OMIT
