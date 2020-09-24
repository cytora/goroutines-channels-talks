package main

import (
	"fmt"
	"time"
)

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

//START OMIT
func main() {
	hat := RabbitHat{contents: "nothing"} //the hat begins by being empty

	fmt.Printf("The magic show has started!\nThe hat contains %s.\n", hat.contents)
	time.Sleep(1 * time.Second)

	fmt.Println("Placing a rabbit in the hat.")
	hatTrickRabbit(hat) // try to place a rabbit in the hat
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
