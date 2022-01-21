package list

type DoubleList struct {
	root BiNode
	free func(interface{})
}

func NewDoubleList(free func(interface{})) *DoubleList {
	if free == nil {
		free = dummy_free
	}
	lst := &DoubleList{free: free}
	lst.root.next = &lst.root
	lst.root.prev = &lst.root
	return lst
}

func (lst *DoubleList) Move() *DoubleList {
	new_lst := &DoubleList{root: lst.root, free: lst.free}

	lst.root.next = &lst.root
	lst.root.prev = &lst.root
	lst.root.val = nil

	return new_lst
}

func (lst *DoubleList) IsEmpty() bool {
	return lst.root.next == &lst.root
}

func (lst *DoubleList) Clear() {
	for !lst.IsEmpty() {
		val, ok := lst.remove(lst.root.next)
		if !ok {
			break
		}
		if val != nil {
			lst.free(val)
		}
	}
}

func (lst *DoubleList) remove(cur *BiNode) (interface{}, bool) {
	if cur == &lst.root {
		return nil, false
	}
	val := cur.Remove()
	return val, true
}

func (lst *DoubleList) PopFront() (interface{}, bool) {
	return lst.remove(lst.root.next)
}

func (lst *DoubleList) PopBack() (interface{}, bool) {
	return lst.remove(lst.root.prev)
}

func (lst *DoubleList) PushFront(val interface{}) *BiNode {
	return lst.root.InsertNext(val)
}

func (lst *DoubleList) PushBack(val interface{}) *BiNode {
	return lst.root.InsertPrev(val)
}

func (lst *DoubleList) Front() Iterator {
	return AsForwardIterator(lst.root.next, &lst.root)
}

func (lst *DoubleList) Back() Iterator {
	return AsReverseIterator(lst.root.prev, &lst.root)
}
