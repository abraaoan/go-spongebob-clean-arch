package cache

import (
	"sync"
	"time"

	"github.com/abraaoan/go-spongebob-clean-arch/internal/entity"
)

type EpisodeCache struct {
	mu       sync.RWMutex
	data     map[string]cachedItem[*entity.Episode]
	lifetime time.Duration
}

func NewEpisodeCache(ttl time.Duration) *EpisodeCache {
	c := &EpisodeCache{
		data:     make(map[string]cachedItem[*entity.Episode]),
		lifetime: ttl,
	}
	go c.startCleanup()
	return c
}

func (c *EpisodeCache) Set(id string, value *entity.Episode) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[id] = cachedItem[*entity.Episode]{
		value,
		time.Now().Add(c.lifetime),
	}
}

func (c *EpisodeCache) Get(id string) (*entity.Episode, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	item, ok := c.data[id]
	if !ok || time.Now().After(item.expiresAt) {
		return nil, false
	}
	return item.value, true
}

func (c *EpisodeCache) startCleanup() {
	ticker := time.NewTicker(c.lifetime)
	for range ticker.C {
		c.cleanup()
	}
}

func (c *EpisodeCache) cleanup() {
	c.mu.Lock()
	defer c.mu.Unlock()
	now := time.Now()
	for k, v := range c.data {
		if now.After(v.expiresAt) {
			delete(c.data, k)
		}
	}
}
