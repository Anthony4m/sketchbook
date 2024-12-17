package main

import (
	"fmt"
	"sketchbook/Page_Replacement"
)

func main() {
	//pages := []int{1, 2, 3, 4, 1, 2, 5, 1, 2, 3, 4, 5}
	////frames := 3
	//lru := NewLRUCache(2)
	//lru.Put(0, pages[0])
	//lru.Put(1, pages[6])
	//lru.Put(2, pages[9])
	//lru.Put(2, pages[0])
	//lru.Get(4)
	//lru.CacheContents()

	page1 := Page_Replacement.NewPage("Page 1")
	page2 := Page_Replacement.NewPage("Page 2")
	page3 := Page_Replacement.NewPage("Page 3")
	//
	//clock := NewClock(2)
	//clock.Put(page1)
	//clock.Put(page2)
	//clock.Put(page3)

	// Create an augmented hash map with capacity 5
	ahm := Page_Replacement.NewAugmentedHashMap(2)

	// Simulate page insertions
	//pages := []int{1, 2, 3, 4, 5, 6, 7, 2, 1, 8}
	pages := []*Page_Replacement.Page{page1, page2, page3, page2, page2, page1}

	for _, page := range pages {
		ahm.Insert(page)
	}

	// Print final state
	ahm.Print()

	// Demonstrate page retrieval
	found, turn := ahm.Get(page2)
	if found {
		fmt.Printf("\nPage 2 found with turn value: %d\n", turn)
	}
}
