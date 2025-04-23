package cache

import (
	"sync"
	"time"

	"github.com/abraaoan/go-spongebob-clean-arch/internal/entity"
)

type CharacterCache struct {
	mu       sync.RWMutex
	data     map[string]cachedItem[*entity.Character]
	lifetime time.Duration
}

type cachedItem[T any] struct {
	value     T
	expiresAt time.Time
}

func NewCharacterCache(ttl time.Duration) *CharacterCache {
	c := &CharacterCache{
		data:     make(map[string]cachedItem[*entity.Character]),
		lifetime: ttl,
	}
	go c.startCleanup()
	return c
}

func (c *CharacterCache) Set(id string, value *entity.Character) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[id] = cachedItem[*entity.Character]{
		value:     value,
		expiresAt: time.Now().Add(c.lifetime),
	}
}

func (c *CharacterCache) Get(id string) (*entity.Character, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	item, ok := c.data[id]
	if !ok || time.Now().After(item.expiresAt) {
		return nil, false
	}
	return item.value, true
}

func (c *CharacterCache) startCleanup() {
	ticker := time.NewTicker(c.lifetime)
	for range ticker.C {
		c.cleanup()
	}
}

func (c *CharacterCache) cleanup() {
	c.mu.Lock()
	defer c.mu.Unlock()
	now := time.Now()
	for k, v := range c.data {
		if now.After(v.expiresAt) {
			delete(c.data, k)
		}
	}
}
