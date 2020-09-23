package main

import (
	"fmt"
	"time"
)

//START OMIT
type Hat struct {
	contents string
}
func hatTrickCopy(hat Hat)       { hat.contents = "a rabbit"}
func hatTrickReference(hat *Hat) { hat.contents = "a rabbit"}
func (h Hat) trickCopy() { h.contents = "flowers" }
func (h *Hat) trickReference() { h.contents = "flowers" }
func main() {
	hat := Hat{contents: "nothing"} //the hat begins by being empty
	fmt.Printf("The magic show has started!\nThe hat contains %s.\n", hat.contents)
	time.Sleep(1 * time.Second)

	hatTrickCopy(hat) //attempt the hat trick without pointers
	fmt.Printf("The hat contains %s.\n", hat.contents)
	time.Sleep(2 * time.Second)
	hatTrickReference(&hat) //attempt the hat trick with pointers
	fmt.Printf("The hat contains %s.\n", hat.contents)
	time.Sleep(2 * time.Second)
	hat.trickCopy()
	fmt.Printf("The hat contains %s.\n", hat.contents)
	time.Sleep(2 * time.Second)
	hat.trickReference()
	fmt.Printf("The hat contains %s.\n", hat.contents)
}
//END OMIT
