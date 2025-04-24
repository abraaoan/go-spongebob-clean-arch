package cache

import (
	"sync"
	"time"

	"github.com/abraaoan/go-spongebob-clean-arch/internal/entity"
)

type SeasonCache struct {
	mu       sync.RWMutex
	data     map[string]cachedItem[*entity.Season]
	lifetime time.Duration
}

func NewSeasonCache(ttl time.Duration) *SeasonCache {
	c := &SeasonCache{
		data:     make(map[string]cachedItem[*entity.Season]),
		lifetime: ttl,
	}
	go c.startCleanup()
	return c
}

func (c *SeasonCache) Set(id string, value *entity.Season) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[id] = cachedItem[*entity.Season]{value, time.Now().Add(c.lifetime)}
}

func (c *SeasonCache) Get(id string) (*entity.Season, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	item, ok := c.data[id]
	if !ok || time.Now().After(item.expiresAt) {
		return nil, false
	}
	return item.value, true
}

func (c *SeasonCache) startCleanup() {
	ticker := time.NewTicker(c.lifetime)
	for range ticker.C {
		c.cleanup()
	}
}

func (c *SeasonCache) cleanup() {
	c.mu.Lock()
	defer c.mu.Unlock()
	now := time.Now()
	for k, v := range c.data {
		if now.After(v.expiresAt) {
			delete(c.data, k)
		}
	}
}
