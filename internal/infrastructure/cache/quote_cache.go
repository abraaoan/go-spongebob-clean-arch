package cache

import (
	"sync"
	"time"

	"github.com/abraaoan/go-spongebob-clean-arch/internal/entity"
)

type QuoteCache struct {
	mu       sync.RWMutex
	data     map[string]cachedItem[*entity.Quote]
	lifetime time.Duration
}

func NewQuoteCache(ttl time.Duration) *QuoteCache {
	c := &QuoteCache{
		data:     make(map[string]cachedItem[*entity.Quote]),
		lifetime: ttl,
	}
	go c.startCleanup()
	return c
}

func (c *QuoteCache) Set(id string, value *entity.Quote) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[id] = cachedItem[*entity.Quote]{value, time.Now().Add(c.lifetime)}
}

func (c *QuoteCache) Get(id string) (*entity.Quote, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	item, ok := c.data[id]
	if !ok || time.Now().After(item.expiresAt) {
		return nil, false
	}
	return item.value, true
}

func (c *QuoteCache) startCleanup() {
	ticker := time.NewTicker(c.lifetime)
	for range ticker.C {
		c.cleanup()
	}
}

func (c *QuoteCache) cleanup() {
	c.mu.Lock()
	defer c.mu.Unlock()
	now := time.Now()
	for k, v := range c.data {
		if now.After(v.expiresAt) {
			delete(c.data, k)
		}
	}
}
