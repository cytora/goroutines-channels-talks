package main

import (
	"fmt"
	"time"
)

//START OMIT
type CopyHat struct {
	contents string
}
func hatTrickCopy(hat CopyHat) {
	hat.contents = "a rabbit"
}
func (h CopyHat) trickCopy() {
	h.contents = "flowers"
}
func main() {
	hat := CopyHat{ contents: "nothing"}
	fmt.Printf("The magic show has started!\nThe hat contains %s.\n", hat.contents)
	time.Sleep(1 * time.Second)

	hatTrickCopy(hat)
	fmt.Printf("The hat contains %s.\n", hat.contents)
	time.Sleep(2 * time.Second)
	hat.trickCopy()
	fmt.Printf("The hat contains %s.\n", hat.contents)
	time.Sleep(2 * time.Second)

	fmt.Println("The magic show has ended!")
}
//END OMIT
