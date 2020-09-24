package main

import (
	"fmt"
	"time"
)

//START OMIT
type RabbitHat struct {
	contents string
}

func hatTrickRabbit(hat RabbitHat) {
	hat.contents = "a rabbit"
	time.Sleep(2 * time.Second)
	fmt.Printf("The hat contains %s inside function.\n", hat.contents)
}

func (hat RabbitHat) trickFlowers() {
	hat.contents = "flowers"
	time.Sleep(2 * time.Second)
	fmt.Printf("The hat contains %s inside function.\n", hat.contents)
}
//END OMIT
