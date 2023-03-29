package cache

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestCache_SetAndGet(t *testing.T) {
	t.Run("SetAndGet_ExistingKey_Success", func(t *testing.T) {
		t.Parallel()
		c := NewCache[string](time.Second, time.Second, 5)
		val := "test value"
		c.Set("test key", &val)
		if res, ok := c.Get("test key"); !ok || *res != val {
			t.Error("Expected value", val, "but got", *res)
		}
	})

	t.Run("Set_Overflow_CapacityLimitEnforced", func(t *testing.T) {
		t.Parallel()
		c := NewCache[string](time.Second, time.Second, 5)
		for i := 0; i < 10; i++ {
			s := fmt.Sprintf("key %d", i)
			c.Set(s, &s)
		}
		if len(c.cache) != c.totalCapacity {
			t.Errorf("Expected cache size %d, but got %d", c.totalCapacity, len(c.cache))
		}
	})

	t.Run("Get_NonExistingKey_FalseReturned", func(t *testing.T) {
		t.Parallel()
		c := NewCache[string](time.Second, time.Second, 5)
		if _, ok := c.Get("non-existing key"); ok {
			t.Error("Expected false, but got true")
		}
	})

	t.Run("Get_ExpiredKey_FalseReturned", func(t *testing.T) {
		t.Parallel()
		c := NewCache[string](time.Second, time.Second, 5)
		val := "test value"
		c.Set("expired key", &val)
		time.Sleep(time.Second * 2)
		if _, ok := c.Get("expired key"); ok {
			t.Error("Expected false, but got true")
		}
	})
}

func TestCache_Purge(t *testing.T) {
	t.Run("Purge_ExpiredKeys_Success", func(t *testing.T) {
		t.Parallel()
		c := NewCache[string](time.Second, time.Second, 5)
		c.Set("test key 1", strPtr("test value 1"))
		c.Set("test key 2", strPtr("test value 2"))
		c.Set("test key 3", strPtr("test value 3"))
		time.Sleep(time.Second * 2)
		c.Purge()
		if len(c.cache) != 0 {
			t.Error("Expected cache size 0, but got", len(c.cache))
		}
	})

	t.Run("Purge_NotExpiredKeys_Success", func(t *testing.T) {
		t.Parallel()
		c := NewCache[string](time.Second, time.Second, 5)
		c.Set("test key 1", strPtr("test value 1"))
		c.Set("test key 2", strPtr("test value 2"))
		c.Set("test key 3", strPtr("test value 3"))
		c.Purge()
		if len(c.cache) != 3 {
			t.Error("Expected cache size 3, but got", len(c.cache))
		}
	})

	t.Run("Purge_CalledConcurrent_Success", func(t *testing.T) {
		t.Parallel()
		c := NewCache[string](time.Second, time.Second, 5)
		var wg sync.WaitGroup
		wg.Add(2)
		go func() {
			defer wg.Done()
			c.Set("test key 1", strPtr("test value 1"))
			c.Set("test key 2", strPtr("test value 2"))
			c.Set("test key 3", strPtr("test value 3"))
		}()
		go func() {
			defer wg.Done()
			time.Sleep(time.Second * 2)
			c.Purge()
		}()
		wg.Wait()
		if len(c.cache) != 0 {
			t.Error("Expected cache size 0, but got", len(c.cache))
		}
	})
}

func TestCache_MoveToHead(t *testing.T) {
	t.Run("MoveToHead_ExistingNode_Success", func(t *testing.T) {
		t.Parallel()
		c := NewCache[string](time.Second, time.Second, 5)
		c.Set("test key 1", strPtr("test value 1"))
		c.Set("test key 2", strPtr("test value 2"))
		c.Set("test key 3", strPtr("test value 3"))
		node, _ := c.cache["test key 1"]
		c.moveToHead(node)
		if res, _ := c.Get("test key 1"); *res != "test value 1" {
			t.Error("Expected value test value 1, but got", *res)
		}
	})

	t.Run("MoveToHead_NonExistingNode_NothingDone", func(t *testing.T) {
		t.Parallel()
		c := NewCache[string](time.Second, time.Second, 5)
		c.Set("test key 1", strPtr("test value 1"))
		c.Set("test key 2", strPtr("test value 2"))
		node := &Node[string]{key: "test key 3", val: nil, expireAt: time.Now().Add(time.Second)}
		c.moveToHead(node)
		if len(c.cache) != 2 {
			t.Error("Expected cache size 2, but got", len(c.cache))
		}
	})
}

func TestCache_RemoveNode(t *testing.T) {
	t.Run("RemoveNode_ExistingNode_Success", func(t *testing.T) {
		t.Parallel()
		c := NewCache[string](time.Second, time.Second, 5)
		c.Set("test key 1", strPtr("test value 1"))
		c.Set("test key 2", strPtr("test value 2"))
		node, _ := c.cache["test key 1"]
		c.removeNode(node)
		if c.linkLen() != 1 {
			t.Error("Expected cache size 1, but got", len(c.cache))
		}
	})

	t.Run("RemoveNode_NonExistingNode_NothingDone", func(t *testing.T) {
		t.Parallel()
		c := NewCache[string](time.Second, time.Second, 5)
		c.Set("test key 1", strPtr("test value 1"))
		node := &Node[string]{key: "test key 2", val: nil, expireAt: time.Now().Add(time.Second)}
		c.removeNode(node)
		if len(c.cache) != 1 {
			t.Error("Expected cache size 1, but got", len(c.cache))
		}
	})
}

func TestCache_RemoveTail(t *testing.T) {
	t.Run("RemoveTail_ExistingNodes_Success", func(t *testing.T) {
		t.Parallel()
		c := NewCache[string](time.Second, time.Second, 3)
		c.Set("test key 1", strPtr("test value 1"))
		c.Set("test key 2", strPtr("test value 2"))
		c.Set("test key 3", strPtr("test value 3"))
		node := c.removeTail()
		if node.key != "test key 1" {
			t.Error("Expected removed node key to be test key 1, but got", node.key)
		}
		if c.linkLen() != 2 {
			t.Error("Expected cache size 2, but got", len(c.cache))
		}
	})

	t.Run("RemoveTail_NoNodes_ReturnsNil", func(t *testing.T) {
		t.Parallel()
		c := NewCache[string](time.Second, time.Second, 3)
		node := c.removeTail()
		if node != nil {
			t.Error("Expected nil, but found a node")
		}
	})
}

func strPtr(str string) *string {
	return &str
}

func TestCache_Concurrency(t *testing.T) {
	t.Run("Parallel_Get_Set_Success", func(t *testing.T) {
		t.Parallel()
		numRoutines := 20
		numSets := 50
		c := NewCache[string](time.Second, time.Second, numRoutines*numSets)
		var wg sync.WaitGroup
		wg.Add(numRoutines)
		for i := 0; i < numRoutines; i++ {
			go func(i int) {
				defer wg.Done()
				for j := 0; j < numSets; j++ {
					key := fmt.Sprintf("test key %d %d", i, j)
					val := fmt.Sprintf("test value %d %d", i, j)
					c.Set(key, &val)
					res, ok := c.Get(key)
					if !ok {
						t.Errorf("Expected key %s to exist, but it doesn't", key)
					}
					if *res != val {
						t.Errorf("Expected value %s for key %s, but got %s", val, key, *res)
					}
				}
			}(i)
		}
		wg.Wait()
		if len(c.cache) != numRoutines*numSets {
			t.Errorf("Expected cache size %d, but got %d", numRoutines*numSets, len(c.cache))
		}
	})

	t.Run("Parallel_Remove_Success", func(t *testing.T) {
		t.Parallel()
		c := NewCache[string](time.Second, time.Second, 100)
		for i := 0; i < 100; i++ {
			key := fmt.Sprintf("test key %d", i)
			val := fmt.Sprintf("test value %d", i)
			c.Set(key, &val)
		}
		var wg sync.WaitGroup
		wg.Add(3)
		go func() {
			defer wg.Done()
			time.Sleep(time.Millisecond * 500)
			for i := 0; i < 50; i++ {
				key := fmt.Sprintf("test key %d", i)
				c.Remove(key)
			}
		}()
		go func() {
			defer wg.Done()
			time.Sleep(time.Millisecond * 250)
			for i := 50; i < 75; i++ {
				key := fmt.Sprintf("test key %d", i)
				c.Remove(key)
			}
		}()
		go func() {
			defer wg.Done()
			time.Sleep(time.Millisecond * 750)
			for i := 75; i < 100; i++ {
				key := fmt.Sprintf("test key %d", i)
				c.Remove(key)
			}
		}()
		wg.Wait()
		if len(c.cache) != 0 {
			t.Errorf("Expected cache size 0, but got %d", len(c.cache))
		}
	})
}
