package list

import (
	"sync"
)

type BiNode struct {
	prev *BiNode
	next *BiNode
	val  interface{}
}

var binodePool = sync.Pool{
	New: func() interface{} {
		return new(BiNode)
	},
}

func new_binode() *BiNode {
	n := binodePool.Get().(*BiNode)
	n.prev = nil
	n.next = nil
	n.val = nil

	return n
}

func delete_binode(n *BiNode) {
	n.prev = nil
	n.next = nil
	n.val = nil

	binodePool.Put(n)
}

func (cur *BiNode) InsertPrev(val interface{}) *BiNode {
	n := new_binode()
	n.val = val

	n.prev = cur.prev
	n.next = cur

	cur.prev.next = n
	cur.prev = n

	return n
}

func (cur *BiNode) InsertNext(val interface{}) *BiNode {
	n := new_binode()
	n.val = val

	n.prev = cur
	n.next = cur.next

	cur.next.prev = n
	cur.next = n

	return n
}

func (cur *BiNode) Value() interface{} {
	return cur.val
}

func (cur *BiNode) Remove() interface{} {
	cur.next.prev = cur.prev
	cur.prev.next = cur.next
	defer delete_binode(cur)

	return cur.val
}

func AsForwardIterator(b, e *BiNode) Iterator {
	return &BiNodeForwardIterator{cur: b, end: e}
}

func AsReverseIterator(b, e *BiNode) Iterator {
	return &BiNodeReverseIterator{cur: b, end: e}
}

type BiNodeForwardIterator struct {
	cur *BiNode
	end *BiNode
}

func (it *BiNodeForwardIterator) Clone() Iterator {
	dup := *it
	return &dup
}

func (it *BiNodeForwardIterator) Next() {
	it.cur = it.cur.next
}

func (it *BiNodeForwardIterator) RemoveAndNext() {
	rm := it.cur
	it.cur = it.cur.next
	rm.Remove()
}

func (it *BiNodeForwardIterator) Insert(val interface{}) {
	it.cur.InsertNext(val)
}

func (it *BiNodeForwardIterator) IsEnd() bool {
	return it.cur == it.end
}

func (it *BiNodeForwardIterator) Value() interface{} {
	return it.cur.val
}

type BiNodeReverseIterator struct {
	cur *BiNode
	end *BiNode
}

func (it *BiNodeReverseIterator) Clone() Iterator {
	dup := *it
	return &dup
}

func (it *BiNodeReverseIterator) Next() {
	it.cur = it.cur.prev
}

func (it *BiNodeReverseIterator) RemoveAndNext() {
	rm := it.cur
	it.cur = it.cur.prev
	rm.Remove()
}

func (it *BiNodeReverseIterator) Insert(val interface{}) {
	it.cur.InsertPrev(val)
}

func (it *BiNodeReverseIterator) IsEnd() bool {
	return it.cur == it.end
}

func (it *BiNodeReverseIterator) Value() interface{} {
	return it.cur.val
}
