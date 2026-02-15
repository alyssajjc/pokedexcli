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
	entries  map[string]cacheEntry
	lock     sync.Mutex
	interval time.Duration
}

func NewCache(interval time.Duration) *Cache {
	newCache := Cache{make(map[string]cacheEntry), sync.Mutex{}, interval}
	go newCache.reapLoop()
	return &newCache
}

func (c *Cache) Add(key string, val []byte) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.entries[key] = cacheEntry{time.Now(), val}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.lock.Lock()
	defer c.lock.Unlock()
	entry, ok := c.entries[key]
	if !ok {
		return nil, false
	}
	return entry.val, true
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()
	for range ticker.C {
		c.lock.Lock()
		for key, val := range c.entries {
			if time.Since(val.createdAt) > c.interval {
				delete(c.entries, key)
			}
		}
		c.lock.Unlock()
	}
}
