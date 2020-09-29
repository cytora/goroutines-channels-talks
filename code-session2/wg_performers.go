package main

import (
	"fmt"
	"sync"
	"time"
)
//START OMIT
func magician(wg sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Magician's WaitGroup is %p\n,", &wg)
	fmt.Println("The magician catches the bullet fired at him!")
	time.Sleep(500 * time.Millisecond)
}

func assistant(wg sync.WaitGroup) {
	defer wg.Done()

	time.Sleep(500 * time.Millisecond)
	fmt.Printf("Assistant's WaitGroup is %p\n", &wg)
	fmt.Println("The assistant fires a bullet at the magician!")
}
//END OMIT