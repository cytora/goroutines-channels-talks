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
	time.Sleep(2 * time.Second)
	fmt.Printf("The hat contains %s inside function.\n", hat.contents)
}

func (hat *PointerRabbitHat) trickFlowers() {
	hat.contents = "flowers"
	time.Sleep(2 * time.Second)
	fmt.Printf("The hat contains %s inside function.\n", hat.contents)
}
//END OMIT
