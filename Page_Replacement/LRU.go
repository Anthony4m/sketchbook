package Page_Replacement

import (
	"container/list"
	"fmt"
	"sync"
)

type LRUCache struct {
	capacity int
	cache    map[int]*list.Element
	queue    *list.List
	mutex    sync.Mutex
	faults   int
}

type cacheItem struct {
	key   int
	value interface{}
}

func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		cache:    make(map[int]*list.Element),
		queue:    list.New(),
		faults:   0,
	}
}

func (lru *LRUCache) Get(key int) (interface{}, bool) {
	lru.mutex.Lock()
	defer lru.mutex.Unlock()

	if elem, found := lru.cache[key]; found {
		// Move the accessed element to the front (most recently used)
		lru.queue.MoveToFront(elem)
		return elem.Value.(*cacheItem).value, true
	}
	return nil, false
}

func (lru *LRUCache) Put(key int, value interface{}) {
	lru.mutex.Lock()
	defer lru.mutex.Unlock()

	// If key already exists, update and move to front
	if elem, found := lru.cache[key]; found {
		lru.queue.MoveToFront(elem)
		elem.Value.(*cacheItem).value = value
		return
	}

	// If cache is full, remove the least recently used item
	if len(lru.cache) >= lru.capacity {
		// Remove from the back of the queue (least recently used)
		back := lru.queue.Back()
		backItem := back.Value.(*cacheItem)
		delete(lru.cache, backItem.key)
		lru.queue.Remove(back)
	}

	// Add new item to the front of the queue
	newElem := lru.queue.PushFront(&cacheItem{
		key:   key,
		value: value,
	})
	lru.cache[key] = newElem
}

func (lru *LRUCache) CacheContents() {
	for _, i := range lru.cache {
		fmt.Println(i.Value)
	}
}
