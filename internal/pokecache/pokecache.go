package pokecache

import (
	"time"
)

func NewCache(interval time.Duration) Cache {

	c := Cache{
		cacheEntry: make(map[string]cacheEntry),
	}
	go c.reapLoop(interval)
	return c
}

func (c *Cache) Add(key string, val []byte) {

	c.mu.Lock()
	c.cacheEntry[key] = cacheEntry{
		createdAt: time.Now().UTC(),
		val:       val,
	}
	c.mu.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	entry, ok := c.cacheEntry[key]
	c.mu.Unlock()
	return entry.val, ok
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(interval)
	}
}

func (c *Cache) reap(interval time.Duration) {
	timeBeforeInterval := time.Now().UTC().Add(-interval)
	for key, value := range c.cacheEntry {
		if value.createdAt.Before(timeBeforeInterval) {
			delete(c.cacheEntry, key)
		}
	}
}
