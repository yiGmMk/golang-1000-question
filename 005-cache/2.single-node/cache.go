package cache

import (
	"sync"

	"github.com/yiGmMk/golang-1000-question/005-cache/2.single-node/lru"
)

type cache struct {
	mu         sync.Mutex
	lru        *lru.Cache
	cacheBytes int64
}

func (c *cache) add(k string, v ByteView) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.lru == nil {
		c.lru = lru.New(c.cacheBytes, nil)
	}
	c.lru.Add(k, v)
}

// 查询缓存
func (c *cache) get(k string) (ByteView, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.lru == nil {
		return ByteView{}, false
	}
	if v, ok := c.lru.Get(k); ok {
		return v.(ByteView), ok
	}
	return ByteView{}, false
}
