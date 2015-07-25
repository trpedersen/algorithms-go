package main

import (
	"fmt"
	//"os"
	//"strconv"
)

func main(){

	head := NewList()
	head.key = 1

	head.InsertAfter(2).InsertAfter(3)
	printList(head)
}

type ListNode struct {
	key  int
	next *ListNode
}

func NewList() *ListNode {
	z := &ListNode{next:nil}
	z.next = z
	head := &ListNode{next: z}
	return head
}

func (node *ListNode) DeleteNext(){
	node.next = node.next.next
}

func (node *ListNode) InsertAfter(v int) *ListNode{
	newNode := &ListNode{key: v, next: node.next}
	node.next = newNode
	return newNode
}

func printList(list *ListNode){
	ptr := list
	fmt.Printf("[")
	for ; ; {
		fmt.Printf(" %d", ptr.key)
		ptr = ptr.next
		if ptr.next == ptr.next.next{
			break
		}
	}
	fmt.Printf("]\n")
}
