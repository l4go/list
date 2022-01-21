package list

import (
	"sync"
)

type Node struct {
	next *Node
	val  interface{}
}

var nodePool = sync.Pool{
	New: func() interface{} {
		return new(Node)
	},
}

func new_node() *Node {
	n := nodePool.Get().(*Node)
	n.next = nil
	n.val = nil

	return n
}

func delete_node(n *Node) {
	n.next = nil
	n.val = nil

	nodePool.Put(n)
}

func (cur *Node) InsertNext(val interface{}) *Node {
	n := new_node()
	n.val = val
	n.next = cur.next

	cur.next = n

	return n
}

func (cur *Node) RemoveNext() (interface{}, bool) {
	if cur.next == nil {
		return nil, false
	}

	del := cur.next
	cur.next = del.next
	defer delete_node(del)

	return del.val, true
}

func (cur *Node) Value() interface{} {
	return cur.val
}

func AsIterator(prev *Node) Iterator {
	return &NodeIterator{prev: prev}
}

type NodeIterator struct {
	prev *Node
}

func (it *NodeIterator) Clone() Iterator {
	dup := *it
	return &dup
}

func (it *NodeIterator) Next() {
	it.prev = it.prev.next
}

func (it *NodeIterator) RemoveAndNext() {
	if it.prev.next == nil {
		return
	}
	it.prev.RemoveNext()
}

func (it *NodeIterator) Insert(val interface{}) {
	it.prev.InsertNext(val)
}

func (it *NodeIterator) IsEnd() bool {
	return it.prev.next == nil
}

func (it *NodeIterator) Value() interface{} {
	return it.prev.next.val
}
