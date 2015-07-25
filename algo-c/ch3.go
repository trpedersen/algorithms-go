package main

import (
	"fmt"
	"os"
	"strconv"
	"github.com/trpedersen/io"
)

func main(){

	stdIn := io.NewStdIn()

	fmt.Println(stdIn)

	node := NewList()
	node.key = 1

	node.InsertAfter(2).InsertAfter(3)
	printList(node)
}

func main1() {

	var args = os.Args[1:]
	if len(args) == 0 {
		fmt.Println("args!")
		os.Exit(1)
	}

	n, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("dodgy!")
		os.Exit(1)
	}

	primes := sieve(n)
	for _, prime := range primes {
		fmt.Printf("%5d", prime)
	}
	fmt.Println()
}

func sieve(n int) []int {

	result := make([]int, 0, 100)
	var a = make([]bool, n+1)
	if n == 0 {
		return result
	}

	a[0] = false
	for i := 1; i <= n; i++ {
		a[i] = true
	}
	for i := 2; i <= n/2; i++ {
		for j := 2; j <= n/i; j++ {
			a[i*j] = false
		}
	}

	for i, isPrime := range a {
		if isPrime {
			result = append(result, i)
		}
	}
	return result

}

type ListNode struct {
	key  int
	next *ListNode
}

func NewList() *ListNode {
	z := &ListNode{}
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
	for ; ptr.next != ptr.next.next; {
		fmt.Printf(" %d", ptr.key)
		ptr = ptr.next
	}
	fmt.Printf("]\n")
}
