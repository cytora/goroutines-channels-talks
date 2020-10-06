package main

import "fmt"

//START OMIT
type InfiniteHat struct {
	contents string
}
func (h *InfiniteHat) placeItem(s string) {
	h.contents = fmt.Sprintf("%s\n%s", h.contents, s)
}
//END OMIT
