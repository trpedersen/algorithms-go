package adt

import (
	"testing"
)

func TestStackPushPopOnce(t *testing.T) {

	s := NewListStack()
	expected := "hello"
	s.Push(expected)
	got := s.Pop().(string)
	if got != expected {
		t.Errorf("stack.Pop(), expected: %s, got: %s\n", expected, got)
	}

}

func TestListStackPeek(t *testing.T) {

	s := NewListStack()
	expected := "hello"
	s.Push(expected)
	got := s.Peek().(string)
	if got != expected {
		t.Errorf("stack.Peek(), expected: %s, got: %s\n", expected, got)
	}

}

func TestListStackPushPopLots(t *testing.T) {

	s := NewListStack()
	for i := 0; i < 1000000; i++ {
		s.Push(i)
	}
	var got int
	for i := 999999; i >= 0; i-- {
		got = s.Pop().(int)
		if got != i {
			t.Errorf("stack.Pop(), expected: %d, got: %d\n", i, got)
		}
	}

}

func BenchmarkListStackPushPop(b *testing.B) {

	s := NewListStack()
	for i := 0; i < b.N; i++ {
		s.Push(i)
	}
	var got int
	for i := b.N - 1; i >= 0; i-- {
		got = s.Pop().(int)
		if got != i {
			b.Errorf("stack.Pop(), expected: %d, got: %d\n", i, got)
		}
	}

}
