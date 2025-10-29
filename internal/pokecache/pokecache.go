package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	mu      sync.Mutex
	entries map[string]cacheEntry
}

func NewCache(interval time.Duration) *Cache {
	spawnedCashe := Cache{
		mu:      sync.Mutex{},
		entries: make(map[string](cacheEntry)),
	}
	spawnedCashe.reapLoop(interval)
	return &spawnedCashe
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.entries[key] = cacheEntry{time.Now(), val}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	result := c.entries[key]
	if _, prs := c.entries[key]; !prs {
		return []byte{}, false
	}
	return (result).val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	go func() {
		for {
			time.Sleep(interval)
			c.mu.Lock()
			for key, val := range c.entries {
				if time.Since(val.createdAt) > interval {
					delete(c.entries, key)
				}
			}
			c.mu.Unlock()
		}
	}()
}
