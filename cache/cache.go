package cache

import (
	"sync"
	"time"
)

// In-memory cache
type Cache struct {
	sync.RWMutex
	values  map[string]interface{}
	expires map[string]time.Time
}

// NewCache creates a new Cache
func NewCache() *Cache {
	return &Cache{
		values:  make(map[string]interface{}),
		expires: make(map[string]time.Time),
	}
}

// Set sets a key-value pair
func (c *Cache) Set(key string, value interface{}, expire time.Duration) {
	c.Lock()
	defer c.Unlock()
	c.values[key] = value
	if expire > 0 {
		c.expires[key] = time.Now().Add(expire)
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
	// if expire, delete the key-value pair
	if expire.Before(time.Now()) {
		c.Lock()
		delete(c.values, key)
		delete(c.expires, key)
		c.Unlock()
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
	c.expires = make(map[string]time.Time)
}

var DBCache = NewCache()

func UseCache(key string, expire time.Duration, fn func() (interface{}, error)) (interface{}, error) {
	// try to get value from cache
	value, ok := DBCache.Get(key)
	if ok {
		return value, nil
	}
	// get value from function
	value, err := fn()

	// set value to cache
	if err != nil {
		return nil, err
	}
	DBCache.Set(key, value, expire)

	return value, nil
}
