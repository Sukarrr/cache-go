package lfu_cache

import (
	"fmt"
	"github.com/Sukarrr/cache-go/node"
	"strconv"
	"testing"
)

func TestLFUCache(t *testing.T) {
	lfu := NewLFUCache(3)
	printLFUCache(lfu)

	for i := 0; i < 4; i++ {
		lfu.Put(strconv.Itoa(i), string(rune('a'+i)))
	}
	printLFUCache(lfu)
	v, _, _ := lfu.Get("3")
	println("Get 3: ", v.(string))
	printLFUCache(lfu)
}

func printLFUCache(c *LFUCache) {
	if c == nil {
		println("nil")
		return
	}
	t := c.minTimes
	if t == 0 {
		return
	}
	cur := c.timesMap[t].L
	count := 0
	for {
		if cur.Pre().(*node.LFUNode) == nil {
			print(t, ": L ")
			cur = cur.Next().(*node.LFUNode)
		} else if cur.Next().(*node.LFUNode) == nil {
			print("R\n")
			t++
			for q := c.timesMap[t]; q == nil; {
				t++
			}
			cur = c.timesMap[t].L
		} else {
			fmt.Printf("%+v ", cur.Value())
			cur = cur.Next().(*node.LFUNode)
			count++
		}
		if count == len(c.nodeMap) {
			break
		}
	}
}
