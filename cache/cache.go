package cache

import (
	"sync"
)

// In-memory cache
type Cache struct {
	sync.RWMutex
	values  map[string]interface{}
	expires map[string]int
}

// NewCache creates a new Cache
func NewCache() *Cache {
	return &Cache{
		values:  make(map[string]interface{}),
		expires: make(map[string]int),
	}
}

// Set sets a key-value pair
func (c *Cache) Set(key string, value interface{}, expire int) {
	c.Lock()
	defer c.Unlock()
	c.values[key] = value
	if expire > 0 {
		c.expires[key] = expire
	}
}

// Get gets a value by key
func (c *Cache) Get(key string) (interface{}, bool) {
	c.RLock()
	defer c.RUnlock()
	value, ok := c.values[key]
	if !ok {
		return nil, false
	}
	expire, ok := c.expires[key]
	if !ok {
		return value, true
	}
	if expire > 0 {
		delete(c.values, key)
		delete(c.expires, key)
		return nil, false
	}
	return value, true
}

// Delete deletes a key-value pair
func (c *Cache) Delete(key string) {
	c.Lock()
	defer c.Unlock()
	delete(c.values, key)
	delete(c.expires, key)
}

// Clear clears all key-value pairs
func (c *Cache) Clear() {
	c.Lock()
	defer c.Unlock()
	c.values = make(map[string]interface{})
	c.expires = make(map[string]int)
}
