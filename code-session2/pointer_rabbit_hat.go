package main

import (
	"fmt"
	"time"
)

//START OMIT
type PointerRabbitHat struct {
	contents string
}

func pointerHatTrickRabbit(hat *PointerRabbitHat) {
	hat.contents = "a rabbit"
	time.Sleep(1 * time.Second)
	fmt.Printf("The Rabbit hat is located at %p\n", hat)
	time.Sleep(1 * time.Second)
	fmt.Printf("The Rabbit hat contains %s inside function.\n", hat.contents)
}

func (hat *PointerRabbitHat) trickFlowers() {
	hat.contents = "flowers"
	time.Sleep(1 * time.Second)
	fmt.Printf("The Flower hat is located at %p\n", hat)
	time.Sleep(1 * time.Second)
	fmt.Printf("The Flower hat contains %s inside function.\n", hat.contents)
}
//END OMIT
