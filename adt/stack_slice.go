package adt

import (
//"errors"
)

type slice_stack []interface{}

func NewSliceStack() Stack {
	slice := slice_stack(make([]interface{}, 0, 0))
	return &slice
}

func (s *slice_stack) Push(item interface{}) {
	*s = append(*s, item)
}

func (s *slice_stack) Pop() (item interface{}) {
	head := len(*s) - 1
	item, *s = (*s)[head], (*s)[:head]
	return
}

func (s *slice_stack) Peek() (item interface{}) {
	if len(*s) > 0 {
		item = (*s)[0]
	}
	return
}

func (s *slice_stack) Size() int {
	return len(*s)
}

func (s *slice_stack) Empty() bool {
	return len(*s) == 0
}
