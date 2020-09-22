package main

import (
	"fmt"
	"sync"
)

//START OMIT
type MutexInfiniteHat struct {
	m *sync.Mutex
	contents string
}

func (h *MutexInfiniteHat) placeItem(s string) {
	h.m.Lock() //acquire control of the hat
	h.contents = fmt.Sprintf("%s\n%s", h.contents, s)
	h.m.Unlock() //relinquish control of the hat once the item is placed
}
//END OMIT