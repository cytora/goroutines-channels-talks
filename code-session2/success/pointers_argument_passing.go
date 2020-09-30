package main

import (
	"fmt"
	"time"
)

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

//START OMIT
func main() {
	hat := PointerRabbitHat{contents: "nothing"} //the hat begins by being empty
	fmt.Println("The magic show has started!")
	fmt.Printf("The original hat is located at %p\nThe hat contains %s.\n", &hat, hat.contents)

	time.Sleep(2 * time.Second)
	fmt.Println("Placing a rabbit in the hat.")
	pointerHatTrickRabbit(&hat) // try to place a rabbit in the hat
	time.Sleep(2 * time.Second)
	fmt.Printf("The original hat contains %s after function.\n", hat.contents)

	time.Sleep(2 * time.Second)
	fmt.Println("Placing flowers in the hat.")
	hat.trickFlowers()
	time.Sleep(2 * time.Second)
	fmt.Printf("The original hat contains %s after function.\n", hat.contents)

	time.Sleep(2 * time.Second)
	fmt.Println("The magic show has ended!")
}
//END OMIT
