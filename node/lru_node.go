package node

import "time"

type LRUNode struct {
	ActiveAt time.Time
	key      string
	value    interface{}
	pre      *LRUNode
	next     *LRUNode
}

func (L *LRUNode) SetKey(s string) {
	L.key = s
}

func (L *LRUNode) SetValue(v interface{}) {
	L.value = v
}

func (L *LRUNode) Key() string {
	return L.key
}

func (L *LRUNode) Value() interface{} {
	return L.value
}

func (L *LRUNode) Pre() Node {
	return L.pre
}

func (L *LRUNode) Next() Node {
	return L.next
}

func (L *LRUNode) SetPre(node Node) {
	L.pre, _ = node.(*LRUNode)
}

func (L *LRUNode) SetNext(node Node) {
	L.next, _ = node.(*LRUNode)
}
