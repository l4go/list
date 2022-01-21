package main

import (
	"log"

	"github.com/l4go/list"
	"strconv"
)

func main() {
	log.Println("---- begin ----")
	defer log.Println("---- end ----")

	lst := list.NewList(nil)
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
	for {
		v, ok := lst.PopFront()
		if !ok {
			break
		}
		log.Printf("POP>%#v\n", v)
	}

	lst.PushFront(5)
	lst.PushFront(6)
	lst.PushFront(7)
	for it := lst.Front(); !it.IsEnd(); it.Next() {
		log.Println("for1:", it.Value())
	}
	for it := lst.Front(); !it.IsEnd(); it.RemoveAndNext() {
		log.Println("remove:", it.Value())
	}

	lst.PushFront(8)
	for it := lst.Front(); !it.IsEnd(); it.Next() {
		log.Println("for2:", it.Value())
	}
}
