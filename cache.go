package explorer

import (
	"sync"
	"time"
)

// Cache is a simple TTL cache to avoid hammering the node with repeated RPC calls.
type Cache struct {
	mu      sync.RWMutex
	entries map[string]cacheEntry
}

type cacheEntry struct {
	value   interface{}
	expires time.Time
}

func NewCache() *Cache {
	c := &Cache{entries: make(map[string]cacheEntry)}
	go c.janitor()
	return c
}

func (c *Cache) Get(key string) (interface{}, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	e, ok := c.entries[key]
	if !ok || time.Now().After(e.expires) {
		return nil, false
	}
	return e.value, true
}

func (c *Cache) Set(key string, value interface{}, ttl time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.entries[key] = cacheEntry{value: value, expires: time.Now().Add(ttl)}
}

func (c *Cache) janitor() {
	for range time.Tick(30 * time.Second) {
		c.mu.Lock()
		now := time.Now()
		for k, v := range c.entries {
			if now.After(v.expires) {
				delete(c.entries, k)
			}
		}
		c.mu.Unlock()
	}
}
