package main

import (
	"log"

	"github.com/l4go/list"
	"strconv"
)

func main() {
	log.Println("---- begin ----")
	defer log.Println("---- end ----")

	lst := list.NewDoubleList(nil)
	defer lst.Clear()

	lst.PushFront(1)
	lst.PushFront(2)
	lst.PushFront(3)
	lst.PushFront(4)
	for it := lst.Front(); !it.IsEnd(); it.Next() {
		log.Println("copy string:", it.Value())
		it.Insert(strconv.Itoa(it.Value().(int)))
		it.Next()
	}
	for it := lst.Back(); !it.IsEnd(); it.Next() {
		log.Println("add 10:", it.Value())
		if v, ok := it.Value().(int); ok {
			it.Insert(v + 10)
			it.Next()
		}
	}
	for {
		v, ok := lst.PopFront()
		if !ok {
			break
		}
		log.Printf("POP>%#v\n", v)
	}

	lst.PushBack(5)
	lst.PushBack(6)
	lst.PushBack(7)
	for it := lst.Front(); !it.IsEnd(); it.Next() {
		log.Println("forward:", it.Value())
	}

	for it := lst.Back(); !it.IsEnd(); it.Next() {
		log.Println("reverse:", it.Value())
	}
	for it := lst.Front(); !it.IsEnd(); it.RemoveAndNext() {
		log.Println("remove:", it.Value())
	}

	lst.PushBack(8)
	for it := lst.Front(); !it.IsEnd(); it.Next() {
		log.Println("for:", it.Value())
	}
}
