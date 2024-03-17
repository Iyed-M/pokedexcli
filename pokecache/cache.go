package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	Entries map[string]cacheEntry
	mu      sync.Mutex
}

type cacheEntry struct {
	val      []byte
	createAt time.Time
}

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		Entries: make(map[string]cacheEntry),
	}
	go c.reapLoop(interval)
	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.Entries[key] = cacheEntry{
		val:      val,
		createAt: time.Now(),
	}
}

func (c *Cache) Get(key string) (val []byte, isFound bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry, ok := c.Entries[key]
	if !ok {
		return nil, false
	}

	return entry.val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for {
		t := <-ticker.C
		c.mu.Lock()
		for key, entry := range c.Entries {
			if t.Sub(entry.createAt) > interval {
				delete(c.Entries, key)
			}
		}
	}
}
