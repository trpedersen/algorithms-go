package main

import (
	"bitbucket.org/trpedersen/algorithms/io"
	"fmt"
	"log"
	"math"
	"os"
)

var parent []int
var sz []int

func main() {

	in, err := os.Open(os.Args[1])
	if err != nil {
		log.Printf("error: %s", err)
		os.Exit(1)
	}
	s := io.NewScanner(in)
	size, err := s.ReadInt()
	if err != nil {
		panic(err)
	}

	log.Printf("N: %d, Log2(N): %f\n", size, math.Log2(float64(size)))

	parent = make([]int, size, size)
	sz = make([]int, size, size)

	for i := 0; i < len(parent); i++ {
		parent[i] = i
		sz[i] = 1
	}

	line, err := s.ReadLine()
	var p, q int
	// count := 0
	for err == nil {
		fmt.Sscanf(line, "%d %d", &p, &q)
		if !connected(p, q) {
			union(p, q)
		}
		line, err = s.ReadLine()
	}

	stdIn := io.NewScanner(os.Stdin)

	//fmt.Println(parent)

	m := make(map[int]int)
	for _, v := range parent {
		m[v] += 1
	}

	fmt.Println(m)
	for {
		p, _ := stdIn.ReadInt()
		q, _ := stdIn.ReadInt()
		fmt.Println(connected(p, q))
	}

}

func find(p int) int {
	// count := 0
	root := p
	for parent[root] != root {
		root = parent[root]
	}
	for p != root {
		q := parent[p]
		parent[p] = root
		p = q
	}
	return root
}

func union(p, q int) {
	idp := find(p)
	idq := find(q)
	if idp != idq {
		if sz[idp] < sz[idq] {
			parent[idp] = idq
			sz[idq] += sz[idp]
		} else {
			parent[idq] = idp
			sz[idp] += sz[idq]
		}
	}
}

func connected(p, q int) bool {
	return find(p) == find(q)
}
