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
	cache map[string]cacheEntry
	mu    *sync.Mutex
}

func (c *Cache) reapLoop(interval time.Duration) {
	t := time.NewTicker(interval)
	for range t.C {
		c.mu.Lock()
		for k := range c.cache {
			if time.Since(c.cache[k].createdAt) > interval {
				delete(c.cache, k)
			}
		}
		c.mu.Unlock()
	}
}

func NewCache(interval time.Duration) Cache {

	c := Cache{cache: make(map[string]cacheEntry),
		mu: &sync.Mutex{}}
	go c.reapLoop(interval)

	return c
}

func (c *Cache) Add(key string, value []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry := cacheEntry{
		createdAt: time.Now(),
		val:       value,
	}

	c.cache[key] = entry
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	value, ok := c.cache[key]
	return value.val, ok
}
