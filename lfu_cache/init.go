package lfu_cache

import (
	"github.com/Sukarrr/cache-go/node"
	"sync"
)

type LFUCache struct {
	nodeMap  map[string]*node.LFUNode
	timesMap map[int]*Queue
	minTimes int
	capacity int
	sync.RWMutex
}

type Queue struct {
	L *node.LFUNode
	R *node.LFUNode
}

func NewLFUCache(capacity int) *LFUCache {
	nodeMap := make(map[string]*node.LFUNode)
	timesMap := make(map[int]*Queue)
	return &LFUCache{
		nodeMap:  nodeMap,
		timesMap: timesMap,
		capacity: capacity,
	}
}

func (c *LFUCache) Put(key string, value interface{}) error {
	if _, ok, err := c.Get(key); !ok && err == nil {
		newNode := &node.LFUNode{
			Times: 1,
		}
		newNode.SetKey(key)
		newNode.SetValue(value)
		c.Lock()
		defer c.Unlock()
		if len(c.nodeMap) >= c.capacity {
			deleteNode := c.timesMap[c.minTimes].L.Next()
			if err = node.DeleteNode(deleteNode); err != nil {
				return err
			} else {
				delete(c.nodeMap, deleteNode.Key())
			}
		}
		c.minTimes = 1
		if t := c.timesMap[newNode.Times]; t == nil {
			newQueue := &Queue{
				L: &node.LFUNode{},
				R: &node.LFUNode{},
			}
			newQueue.L.SetNext(newQueue.R)
			newQueue.R.SetPre(newQueue.L)
			c.timesMap[newNode.Times] = newQueue
		}
		if err = node.AddNode(c.timesMap[newNode.Times].R.Pre(), newNode); err != nil {
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

func (c *LFUCache) Get(key string) (value interface{}, ok bool, err error) {
	c.Lock()
	defer c.Unlock()
	var v *node.LFUNode
	if v, ok = c.nodeMap[key]; !ok {
		return
	} else {
		copyV := &node.LFUNode{
			Times: v.Times + 1,
		}
		copyV.SetKey(key)
		copyV.SetValue(v.Value())
		if err = node.DeleteNode(v); err != nil {
			return
		}
		if t := c.timesMap[c.minTimes]; t != nil && t.L != nil && t.L.Next() == t.R {
			c.minTimes++
		}
		if t := c.timesMap[copyV.Times]; t == nil {
			newQueue := &Queue{
				L: &node.LFUNode{},
				R: &node.LFUNode{},
			}
			newQueue.L.SetNext(newQueue.R)
			newQueue.R.SetPre(newQueue.L)
			c.timesMap[copyV.Times] = newQueue
		}
		if err = node.AddNode(c.timesMap[copyV.Times].R.Pre(), copyV); err != nil {
			return
		}

		c.nodeMap[key] = copyV
		return copyV.Value(), ok, nil
	}
}
