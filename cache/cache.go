package cache

import (
	"sync"
	"time"

	"github.com/wao3/luogu-stats-card/model/fetch"
)

const DefaultExpireTime = 12 * time.Hour
const DefaultPurgeTime = 5 * time.Minute

type Cache[T fetch.DataType] struct {
	expireTime time.Duration
	purgeTime  time.Duration
	mu         sync.RWMutex
	cache      map[string]cacheValue[T]
}

type cacheValue[T fetch.DataType] struct {
	data     *T
	expireAt time.Time
}

func NewCache[T fetch.DataType](expireTime, purgeTime time.Duration) *Cache[T] {
	c := &Cache[T]{
		expireTime: expireTime,
		purgeTime:  purgeTime,
		cache:      make(map[string]cacheValue[T]),
	}
	go func() {
		for {
			time.Sleep(purgeTime)
			c.Purge()
		}
	}()
	return c
}

func (c *Cache[T]) Set(key string, data *T) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cache[key] = cacheValue[T]{
		data:     data,
		expireAt: time.Now().Add(c.expireTime),
	}
}

func (c *Cache[T]) Get(key string) (*T, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	value, ok := c.cache[key]
	if !ok {
		return nil, false
	}
	if value.expireAt.Before(time.Now()) {
		return nil, false
	}
	return value.data, true
}

func (c *Cache[T]) Purge() {
	c.mu.Lock()
	defer c.mu.Unlock()
	for key, value := range c.cache {
		if value.expireAt.Before(time.Now()) {
			delete(c.cache, key)
		}
	}
}
