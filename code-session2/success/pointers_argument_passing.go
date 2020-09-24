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
	time.Sleep(2 * time.Second)
	fmt.Printf("The hat contains %s inside function.\n", hat.contents)
}

func (hat *PointerRabbitHat) trickFlowers() {
	hat.contents = "flowers"
	time.Sleep(2 * time.Second)
	fmt.Printf("The hat contains %s inside function.\n", hat.contents)
}

//START OMIT
func main() {
	hat := PointerRabbitHat{contents: "nothing"} //the hat begins by being empty

	fmt.Printf("The magic show has started!\nThe hat contains %s.\n", hat.contents)
	time.Sleep(1 * time.Second)

	fmt.Println("Placing a rabbit in the hat.")
	pointerHatTrickRabbit(&hat) //pass a pointer to the hat
	time.Sleep(2 * time.Second)
	fmt.Printf("The hat contains %s after function.\n", hat.contents)

	fmt.Println("Placing flowers in the hat.")
	hat.trickFlowers()
	time.Sleep(2 * time.Second)
	fmt.Printf("The hat contains %s after function.\n", hat.contents)

	time.Sleep(2 * time.Second)
	fmt.Println("The magic show has ended!")
}
//END OMIT
