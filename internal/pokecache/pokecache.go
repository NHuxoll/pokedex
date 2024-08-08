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
	cache    map[string]cacheEntry
	duration time.Duration
	mu       sync.Mutex
}

func NewCache(duration time.Duration) *Cache {
	c := Cache{cache: make(map[string]cacheEntry), duration: duration, mu: sync.Mutex{}}
	c.reap()
	return &c
}

func (c *Cache) Add(key string, value []byte) {
	entry := cacheEntry{
		createdAt: time.Now(),
		val:       value,
	}
	c.cache[key] = entry
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	entry, ok := c.cache[key]
	if !ok {
		return nil, false
	}
	entry.createdAt = time.Now()
	val := entry.val
	c.mu.Unlock()
	return val, true
}

func (c *Cache) reap() {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	done := make(chan struct{})
	go func() {
		time.Sleep(60 * 5 * time.Second)
		done <- true
	}()
	for {
		select {
		case <-done:
			return
		case <-ticker.C:
			c.mu.Lock()
			for key, entry := range c.cache {
				if time.Now()-entry.createdAt > c.duration {
					delete(c.cache, key)
				}
			}
			c.mu.Unlock()
	}
}
