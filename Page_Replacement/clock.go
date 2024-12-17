package Page_Replacement

import (
	"fmt"
	"sync"
)

type clock struct {
	capacity int
	cache    map[string]*cacheItemClock
	hand     int
	mutex    sync.Mutex
	faults   int
}

type cacheItemClock struct {
	key      string
	pageItem *Page
}

func NewClock(capacity int) *clock {
	return &clock{
		cache:    make(map[string]*cacheItemClock),
		capacity: capacity,
		hand:     0,
		faults:   0,
	}
}

func (clk *clock) Get(item *Page) *Page {
	if elem, found := clk.cache[item.id]; found {
		currpage := elem.pageItem
		currpage.reference = 1
		return currpage
	}
	clk.faults += 1
	return nil
}

func (clk *clock) Put(newpage *Page) {
	// If the page is already in the cache, reset its reference
	if elem, found := clk.cache[newpage.id]; found {
		elem.pageItem.reference = 1
		return
	}

	// If the cache is full, evict pages starting from the current hand
	if len(clk.cache) >= clk.capacity {
		keys := make([]string, 0, len(clk.cache))
		for key := range clk.cache {
			keys = append(keys, key) // Collect all keys to simulate "iteration"
		}

		for {
			// Simulate circular traversal using the hand index
			key := keys[clk.hand]
			item := clk.cache[key]
			fmt.Println(item.pageItem)

			if item.pageItem.reference == 0 {
				// Evict this item
				delete(clk.cache, key)
				break
			}

			// Decrement reference count and move the hand forward
			item.pageItem.reference--
			clk.hand = (clk.hand + 1) % len(clk.cache)
		}
	}

	// Add the new page to the cache
	newElem := &cacheItemClock{
		pageItem: newpage,
	}
	clk.cache[newpage.id] = newElem
}
