package node

import (
	"errors"
)

type Node interface {
	Pre() Node
	Next() Node
	SetPre(node Node)
	SetNext(node Node)
	Key() string
	Value() interface{}
	SetKey(s string)
	SetValue(v interface{})
}

func AddNode(root, newNode Node) error {
	if root == nil || root.Next() == nil || newNode == nil {
		return errors.New("root or root.next or newNode is nil")
	}
	next := root.Next()
	root.SetNext(newNode)
	newNode.SetPre(root)
	newNode.SetNext(next)
	next.SetPre(newNode)
	return nil
}

func DeleteNode(root Node) error {
	if root == nil || root.Pre() == nil || root.Next() == nil {
		return errors.New("root or pre or next is nil")
	}
	pre := root.Pre()
	next := root.Next()
	pre.SetNext(next)
	next.SetPre(pre)
	return nil
}
