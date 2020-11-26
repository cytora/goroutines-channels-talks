package main

import (
	"fmt"
	"time"
)

//START OMIT
type PointerHat struct {
	contents string
}
func hatTrickReference(hat *PointerHat) {
	hat.contents = "a rabbit"
}
func (h *PointerHat) trickReference() {
	h.contents = "flowers"
}
func main() {
	hat := PointerHat{ contents: "nothing" }
	fmt.Printf("The magic show has started!\nThe hat contains %s.\n", hat.contents)
	time.Sleep(1 * time.Second)

	hatTrickReference(&hat)
	fmt.Printf("The hat contains %s.\n", hat.contents)
	time.Sleep(2 * time.Second)
	hat.trickReference()
	fmt.Printf("The hat contains %s.\n", hat.contents)
	time.Sleep(2 * time.Second)

	fmt.Println("The magic show has ended!")
}
//END OMIT
