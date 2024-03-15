package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	Entries  map[string]cacheEntry
	mu       sync.Mutex
	interval time.Duration
}

type cacheEntry struct {
	val      []byte
	createAt time.Time
}

func NewCache(interval time.Duration) *Cache {
	return &Cache{
		Entries:  make(map[string]cacheEntry),
		interval: interval,
	}
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

func (c *Cache) ReapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
}
