// Implementation based on the algorithm described in:
// "Improve the Performance of LRU Page Replacement Algorithm using Augmentation of Data Structure"
// Authors: Mr. C.C. Kavar, Mr. S. S. Parmar
// Published in: IEEE - 31661, 4th ICCCNT, July 4-6, 2013.
// Reference: https://doi.org/10.1109/ICCCNT.2013.6726496
// Description: This implementation applies the augmented Hash Table algorithm for improved efficiency.

package Page_Replacement

import (
	"fmt"
	"sync"
)

// AugmentedHashMap implements the turn-based LRU tracking
type AugmentedHashMap struct {
	// Use mutex for thread-safety
	mu sync.RWMutex

	// Internal map to store pages
	pages map[*Page]PageEntry

	// Maximum capacity of the map
	capacity int

	// Tracking counters
	frequency   int
	hitCounter  int
	missCounter int
	currentTurn int
}

// PageEntry stores additional metadata for each page
type PageEntry struct {
	Turn int // Turn value for LRU tracking
}

// NewAugmentedHashMap creates a new augmented hash map
func NewAugmentedHashMap(capacity int) *AugmentedHashMap {
	return &AugmentedHashMap{
		pages:       make(map[*Page]PageEntry, capacity),
		capacity:    capacity,
		frequency:   0,
		hitCounter:  0,
		missCounter: 0,
		currentTurn: 0,
	}
}

// Insert handles page insertion with turn-based LRU logic
func (ahm *AugmentedHashMap) Insert(page *Page) {
	ahm.mu.Lock()
	defer ahm.mu.Unlock()

	// Increment turn and frequency
	ahm.currentTurn++
	ahm.frequency++

	// Check if page already exists
	if _, exists := ahm.pages[page]; exists {
		// Page hit
		ahm.pages[page] = PageEntry{Turn: ahm.currentTurn}
		ahm.hitCounter++
		return
	}

	// Page miss
	ahm.missCounter++

	// Check if map is at capacity
	if len(ahm.pages) >= ahm.capacity {
		// Find and remove least recently used page
		var lruPage *Page
		var lruTurn = ahm.currentTurn

		for p, entry := range ahm.pages {
			if entry.Turn < lruTurn {
				lruPage = p
				lruTurn = entry.Turn
			}
		}

		// Remove least recently used page
		delete(ahm.pages, lruPage)
	}

	// Insert new page
	ahm.pages[page] = PageEntry{Turn: ahm.currentTurn}
}

// Get retrieves a page and updates its turn
func (ahm *AugmentedHashMap) Get(page *Page) (bool, int) {
	ahm.mu.Lock()
	defer ahm.mu.Unlock()

	entry, exists := ahm.pages[page]
	if exists {
		// Update turn when page is accessed
		ahm.currentTurn++
		entry.Turn = ahm.currentTurn
		ahm.pages[page] = entry
		return true, entry.Turn
	}

	return false, 0
}

// Print displays the current state of the hash map
func (ahm *AugmentedHashMap) Print() {
	ahm.mu.RLock()
	defer ahm.mu.RUnlock()

	fmt.Println("Augmented Hash Map State:")
	for page, entry := range ahm.pages {
		fmt.Printf("Page: %d, Turn: %d\n", page, entry.Turn)
	}

	fmt.Printf("\nStatistics:\n")
	fmt.Printf("Capacity: %d\n", ahm.capacity)
	fmt.Printf("Current Size: %d\n", len(ahm.pages))
	fmt.Printf("Hit Counter: %d\n", ahm.hitCounter)
	fmt.Printf("Miss Counter: %d\n", ahm.missCounter)
	fmt.Printf("Current Frequency: %d\n", ahm.frequency)
}
