package adt

import (
//"errors"
)

type list_node struct {
	item interface{}
	next *list_node
}

type list_stack struct {
	head *list_node
	size int
}

func NewListStack() Stack {
	return &list_stack{}
}

func (s *list_stack) Push(item interface{}) {
	list_node := &list_node{
		item: item,
		next: s.head,
	}
	s.head = list_node
	s.size += 1
	return
}

func (s *list_stack) Pop() (item interface{}) {
	if s.head != nil {
		item = s.head.item
		s.head = s.head.next
		s.size -= 1
	}
	return
}

func (s *list_stack) Peek() (item interface{}) {
	if s.head != nil {
		return s.head.item
	} else {
		return nil
	}
}

func (s *list_stack) Size() int {
	return s.size
}

func (s *list_stack) Empty() bool {
	return s.head == nil
}
