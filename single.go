package list

type List struct {
	root Node
	free func(interface{})
}

func NewList(free func(interface{})) *List {
	if free == nil {
		free = dummy_free
	}
	lst := &List{free: free}
	return lst
}

func (lst *List) Move() *List {
	new_lst := &List{root: lst.root, free: lst.free}
	lst.root.next = nil
	lst.root.val = nil

	return new_lst
}

func (lst *List) IsEmpty() bool {
	return lst.root.next == nil
}

func (lst *List) Clear() {
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

func (lst *List) PopFront() (interface{}, bool) {
	return lst.root.RemoveNext()
}

func (lst *List) PushFront(val interface{}) *Node {
	return lst.root.InsertNext(val)
}

func (lst *List) Front() Iterator {
	return AsIterator(&lst.root)
}
