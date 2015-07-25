package adt

type Stack interface {
	Push(item interface{})
	Pop() (item interface{})
	Peek() (item interface{})
	Size() int
	Empty() bool
}
