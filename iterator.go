package list

type Value interface {
	Value() interface{}
}

type Iterator interface {
	Clone() Iterator
	Next()
	RemoveAndNext()
	Insert(interface{})
	IsEnd() bool
	Value() interface{}
}
