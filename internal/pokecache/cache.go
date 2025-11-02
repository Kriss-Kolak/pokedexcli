package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	mu       *sync.Mutex
	v        map[string]cacheEntry
	interval time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	new_entry := cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	c.v[key] = new_entry

}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	value, ok := c.v[key]
	if !ok {
		return []byte{}, false
	}
	return value.val, true
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	for range ticker.C {
		c.mu.Lock()
		for k, v := range c.v {
			if time.Since(v.createdAt) > c.interval {
				delete(c.v, k)
			}
		}
		c.mu.Unlock()
	}

}

func NewCache(interval time.Duration) *Cache {
	new_cache := &Cache{
		mu:       &sync.Mutex{},
		v:        map[string]cacheEntry{},
		interval: interval}
	go new_cache.reapLoop()
	return new_cache
}
