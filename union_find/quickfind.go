package main

import (
	"bitbucket.org/trpedersen/algorithms/io"
	"fmt"
	"log"
	"os"
)

var ids []int

func main() {

	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Printf("error: %s", err)
		os.Exit(1)
	}
	s := io.NewScanner(f)
	size, err := s.ReadInt()
	if err != nil {
		panic(err)
	}

	ids = make([]int, size, size)

	for i := 0; i < len(ids); i++ {
		ids[i] = i
	}

	line, err := s.ReadLine()
	var p, q int
	count := 0
	for err == nil {
		fmt.Sscanf(line, "%d %d", &p, &q)
		if !connected(p, q) {
			union(p, q)
		}
		count += 1
		if count%1000 == 0 {
			log.Println(count)
		}
		line, err = s.ReadLine()
	}

	fmt.Println(ids)

}

func union(p, q int) {
	if connected(p, q) {
		return
	}
	idp := ids[p]
	idq := ids[q]
	for i := 0; i < len(ids); i++ {
		if ids[i] == idp {
			ids[i] = idq
		}
	}
}

func find(p int) int {
	return ids[p]
}

func connected(p, q int) bool {
	return ids[p] == ids[q]
}
