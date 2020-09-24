package main

import (
	"fmt"
	"time"
)

//START OMIT
type ImpossibleRabbitHat struct {
	contents string
}

func swapHat(hat *ImpossibleRabbitHat) {
	rabbitHat := ImpossibleRabbitHat{
		contents: "rabbit",
	}
	fmt.Printf("New hat pointer is %p\n", &rabbitHat)
	time.Sleep(2 * time.Second)
	fmt.Printf("New hat contains %s\n", rabbitHat.contents)
	time.Sleep(2 * time.Second)

	hat = &rabbitHat
	fmt.Printf("Hat pointer during swap is %p\n", hat)
	time.Sleep(2 * time.Second)
	fmt.Printf("Hat contents during swap is %s\n", hat.contents)
}
//END OMIT
