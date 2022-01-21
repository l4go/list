package main

import (
	"log"

	"github.com/l4go/list"
	"strconv"
)

func main() {
	log.Println("---- begin ----")
	defer log.Println("---- end ----")

	lst := list.NewFifoList(nil)
	defer lst.Clear()

	lst.PushBack(1)
	lst.PushBack(2)
	lst.PushBack(3)
	lst.PushBack(4)
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

	lst.PushBack(5)
	lst.PushBack(6)
	lst.PushBack(7)
	for it := lst.Front(); !it.IsEnd(); it.Next() {
		log.Println("for1:", it.Value())
	}
	for it := lst.Front(); !it.IsEnd(); it.RemoveAndNext() {
		log.Println("remove:", it.Value())
	}

	lst.PushBack(8)
	for it := lst.Front(); !it.IsEnd(); it.Next() {
		log.Println("for2:", it.Value())
	}
}
