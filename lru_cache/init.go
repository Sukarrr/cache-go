package lru_cache

import (
	"github.com/Sukarrr/cache-go/node"
	"sync"
	"time"
)

type LRUCache struct {
	nodeMap  map[string]*node.LRUNode
	capacity int
	sync.RWMutex
	L *node.LRUNode // delete
	R *node.LRUNode // add
}

func NewLRUCache(capacity int) *LRUCache {
	nodeMap := make(map[string]*node.LRUNode)
	L := &node.LRUNode{}
	R := &node.LRUNode{}
	L.SetNext(R)
	R.SetPre(L)
	return &LRUCache{
		nodeMap:  nodeMap,
		capacity: capacity,
		L:        L,
		R:        R,
	}
}

func (c *LRUCache) Put(key string, value interface{}) error {
	if _, ok, err := c.Get(key); !ok && err == nil {
		newNode := &node.LRUNode{
			ActiveAt: time.Now(),
		}
		newNode.SetKey(key)
		newNode.SetValue(value)
		c.Lock()
		defer c.Unlock()
		if len(c.nodeMap) >= c.capacity {
			deleteKey := c.L.Next().Key()
			if err = node.DeleteNode(c.L.Next()); err != nil {
				return err
			} else {
				delete(c.nodeMap, deleteKey)
			}
		}
		if err = node.AddNode(c.R.Pre(), newNode); err != nil {
			return err
		}
		c.nodeMap[key] = newNode
		return nil
	} else if err != nil {
		return err
	} else {
		c.nodeMap[key].SetValue(value)
	}
	return nil
}

func (c *LRUCache) Get(key string) (value interface{}, ok bool, err error) {
	c.Lock()
	defer c.Unlock()
	var v *node.LRUNode
	if v, ok = c.nodeMap[key]; !ok {
		return
	} else {
		copyV := &node.LRUNode{
			ActiveAt: time.Now(),
		}
		copyV.SetKey(key)
		copyV.SetValue(v.Value())
		if err = node.DeleteNode(v); err != nil {
			return
		}
		if err = node.AddNode(c.R.Pre(), copyV); err != nil {
			return
		}
		c.nodeMap[key] = copyV
		return copyV.Value(), ok, nil
	}
}
