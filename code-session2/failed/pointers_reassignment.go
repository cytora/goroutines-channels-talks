package main

import (
	"fmt"
	"time"
)

type ImpossibleRabbitHat struct {
	contents string
}

func swapHat(hat *ImpossibleRabbitHat) {
	rabbitHat := ImpossibleRabbitHat{
		contents: "rabbit",
	}
	time.Sleep(2 * time.Second)
	fmt.Printf("New hat pointer is %p\n", &rabbitHat)
	fmt.Printf("New hat contains %s\n", rabbitHat.contents)
	time.Sleep(2 * time.Second)

	hat = &rabbitHat
	fmt.Printf("Hat pointer during swap is %p\n", hat)
	time.Sleep(2 * time.Second)
	fmt.Printf("Hat contents during swap is %s\n", hat.contents)
}

//START OMIT
func main() {
	hat := &ImpossibleRabbitHat{contents: "nothing"} //the hat begins by being empty
	fmt.Println("The magic show has started!")

	time.Sleep(2* time.Second)
	fmt.Printf("Hat pointer is %p before swap\n", hat)
	time.Sleep(2* time.Second)
	fmt.Printf("Hat contains %s before swap\n", hat.contents)

	time.Sleep(2* time.Second)
	fmt.Println("Commencing the hat swap")
	swapHat(hat) // send the pointer to swap the hat

	time.Sleep(2 * time.Second)
	fmt.Printf("Hat pointer is %p after swap\n", hat)
	time.Sleep(2* time.Second)
	fmt.Printf("Hat contains %s after swap\n", hat.contents)
	time.Sleep(2* time.Second)
	fmt.Println("The magic show has ended!")
}
//END OMIT
