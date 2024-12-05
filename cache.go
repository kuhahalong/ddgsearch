package ddgsearch

import (
	"sync"
	"time"
)

// cache implements a simple in-memory cache with expiration
type cache struct {
	mu       sync.RWMutex
	items    map[string]*cacheItem
	maxAge   time.Duration
}

type cacheItem struct {
	value      interface{}
	expiration time.Time
}

func newCache(maxAge time.Duration) *cache {
	return &cache{
		items:  make(map[string]*cacheItem),
		maxAge: maxAge,
	}
}

func (c *cache) get(key string) (interface{}, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	item, exists := c.items[key]
	if !exists {
		return nil, false
	}

	if time.Now().After(item.expiration) {
		delete(c.items, key)
		return nil, false
	}

	return item.value, true
}

func (c *cache) set(key string, value interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.items[key] = &cacheItem{
		value:      value,
		expiration: time.Now().Add(c.maxAge),
	}
}

func (c *cache) delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.items, key)
}

func (c *cache) clear() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.items = make(map[string]*cacheItem)
}
