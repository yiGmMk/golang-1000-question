package cache

import (
	"fmt"
	"sync"
)

// 缓存未命中时从数据源取数据的方法
type Getter interface {
	Get(k string) ([]byte, error)
}

// GetterFunc实现Getter接口
var _ Getter = (*GetterFunc)(nil)

type GetterFunc func(k string) ([]byte, error)

func (f GetterFunc) Get(k string) ([]byte, error) {
	return f(k)
}

// Group 可以认为是一个缓存的命名空间，每个 Group 拥有一个唯一的名称 name。比如可以创建三个 Group，缓存学生的成绩命名为 scores，
// 缓存学生信息的命名为 info，缓存学生课程的命名为 courses。
type Group struct {
	name      string
	getter    Getter
	mainCache cache
}

var (
	mu     sync.RWMutex
	groups = make(map[string]*Group)
)

func NewGroup(name string, cacheBytes int64, getter Getter) *Group {
	if getter == nil {
		panic("nil Getter")
	}
	mu.Lock()
	defer mu.Unlock()
	g := &Group{
		name:      name,
		getter:    getter,
		mainCache: cache{cacheBytes: cacheBytes},
	}
	groups[name] = g
	return g
}

func (g *Group) Get(k string) (ByteView, error) {
	if k == "" {
		return ByteView{}, fmt.Errorf("key is required")
	}

	if v, ok := g.mainCache.get(k); ok {
		return v, nil
	}
	return g.load(k)
}

func (g *Group) load(key string) (value ByteView, err error) {
	bytes, err := g.getter.Get(key)
	if err != nil {
		return ByteView{}, err
	}
	value = ByteView{b: cloneBytes(bytes)}
	g.populateCache(key, value)
	return value, nil
}

func (g *Group) populateCache(key string, value ByteView) {
	g.mainCache.add(key, value)
}

// GetGroup returns the named group previously created with NewGroup, or
// nil if there's no such group.
func GetGroup(name string) *Group {
	mu.RLock()
	g := groups[name]
	mu.RUnlock()
	return g
}
