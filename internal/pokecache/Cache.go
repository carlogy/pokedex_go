package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cacheEntry map[string]cacheEntry
	mu         sync.RWMutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}
