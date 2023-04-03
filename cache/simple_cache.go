package cache

import (
	"sync"
	"time"
)

const DefaultExpireTime = 12 * time.Hour
const DefaultPurgeTime = 5 * time.Minute
const DefaultCapacity = 10000

type Node[T any] struct {
	key        string
	val        *T
	expireAt   time.Time
	prev, next *Node[T]
}

type Cache[T any] struct {
	totalCapacity int
	head, tail    *Node[T]
	cache         map[string]*Node[T]
	expireTime    time.Duration
	purgeTime     time.Duration
	mu            sync.RWMutex
}

func NewCache[T any](expireTime, purgeTime time.Duration, capacity int) *Cache[T] {
	head, tail := &Node[T]{}, &Node[T]{}
	head.next, tail.prev = tail, head
	c := &Cache[T]{
		expireTime:    expireTime,
		purgeTime:     purgeTime,
		cache:         make(map[string]*Node[T]),
		totalCapacity: capacity,
		head:          head,
		tail:          tail,
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
	if node, ok := c.cache[key]; ok {
		node.val = data
		node.expireAt = time.Now().Add(c.expireTime)
		c.moveToHead(node)
		return
	}
	node := &Node[T]{key, data, time.Now().Add(c.expireTime), nil, nil}
	c.cache[key] = node
	c.addToHead(node)

	if len(c.cache) > c.totalCapacity {
		removed := c.removeTail()
		delete(c.cache, removed.key)
	}
}

func (c *Cache[T]) Get(key string) (*T, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	node, ok := c.cache[key]
	if !ok {
		return nil, false
	}
	if node.expireAt.Before(time.Now()) {
		return nil, false
	}
	c.moveToHead(node)
	return node.val, true
}

func (c *Cache[T]) Remove(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	node, ok := c.cache[key]
	if !ok {
		return
	}
	c.removeNode(node)
	delete(c.cache, key)
	return
}

func (c *Cache[T]) Purge() {
	c.mu.Lock()
	defer c.mu.Unlock()
	for key, value := range c.cache {
		if value.expireAt.Before(time.Now()) {
			delete(c.cache, key)
			c.removeNode(value)
		}
	}
}

func (c *Cache[T]) addToHead(node *Node[T]) {
	node.next = c.head.next
	node.prev = c.head
	c.head.next.prev = node
	c.head.next = node
}

func (c *Cache[T]) removeNode(node *Node[T]) {
	if _, ok := c.cache[node.key]; !ok {
		return
	}
	if node == c.tail || node == c.head {
		return
	}
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (c *Cache[T]) moveToHead(node *Node[T]) {
	if _, ok := c.cache[node.key]; !ok {
		return
	}
	c.removeNode(node)
	c.addToHead(node)
}

func (c *Cache[T]) removeTail() *Node[T] {
	node := c.tail.prev
	if node == c.head {
		return nil
	}
	c.removeNode(node)
	return node
}

func (c *Cache[T]) linkLen() int {
	count := 0
	for node := c.head.next; node != c.tail; node = node.next {
		count++
	}
	return count
}
