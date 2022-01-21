package list

type FifoList struct {
	root Node
	tail *Node
	free func(interface{})
}

func NewFifoList(free func(interface{})) *FifoList {
	if free == nil {
		free = dummy_free
	}
	lst := &FifoList{free: free}
	lst.tail = &lst.root
	return lst
}

func (lst *FifoList) Move() *FifoList {
	new_lst := &FifoList{root: lst.root, tail: lst.tail, free: lst.free}
	lst.root.next = nil
	lst.root.val = nil
	lst.tail = &lst.root

	return new_lst
}

func (lst *FifoList) IsEmpty() bool {
	return lst.root.next == nil
}

func (lst *FifoList) Clear() {
	for !lst.IsEmpty() {
		val, ok := lst.root.RemoveNext()
		if !ok {
			break
		}
		if val != nil {
			lst.free(val)
		}
	}
}

func (lst *FifoList) PopFront() (interface{}, bool) {
	if lst.root.next == lst.tail {
		lst.tail = &lst.root
	}
	return lst.root.RemoveNext()
}

func (lst *FifoList) PushBack(val interface{}) *Node {
	n := lst.tail.InsertNext(val)
	lst.tail = lst.tail.next
	return n
}

func (lst *FifoList) Front() Iterator {
	return AsIterator(&lst.root)
}
