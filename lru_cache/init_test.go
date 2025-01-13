package lru_cache

import (
	"fmt"
	"github.com/Sukarrr/cache-go/node"
	"strconv"
	"testing"
)

func TestLRUCache(t *testing.T) {
	lru := NewLRUCache(10)
	printLRUCache(lru)

	for i := 0; i < 11; i++ {
		lru.Put(strconv.Itoa(i), string(rune('a'+i)))
	}
	printLRUCache(lru)
	v, _, _ := lru.Get("5")
	println("Get 5: ", v.(string))
	printLRUCache(lru)
}

func printLRUCache(c *LRUCache) {
	if c == nil {
		println("nil")
		return
	}
	for cur := c.L; cur != nil; cur = cur.Next().(*node.LRUNode) {
		if cur == c.L {
			print("L ")
		} else if cur == c.R {
			print("R\n")
		} else {
			fmt.Printf("%+v ", cur.Value())
		}
	}
}
