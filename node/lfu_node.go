package node

type LFUNode struct {
	Times int
	key   string
	value interface{}
	pre   *LFUNode
	next  *LFUNode
}

func (L *LFUNode) SetKey(s string) {
	L.key = s
}

func (L *LFUNode) SetValue(v interface{}) {
	L.value = v
}

func (L *LFUNode) Key() string {
	return L.key
}

func (L *LFUNode) Value() interface{} {
	return L.value
}

func (L *LFUNode) Pre() Node {
	return L.pre
}

func (L *LFUNode) Next() Node {
	return L.next
}

func (L *LFUNode) SetPre(node Node) {
	L.pre, _ = node.(*LFUNode)
}

func (L *LFUNode) SetNext(node Node) {
	L.next, _ = node.(*LFUNode)
}
