package main

import (
	"bitbucket.org/trpedersen/algorithms/io"
	"fmt"
	"os"
	"sort"
)

func main() {

	var (
		scanner   *io.Scanner
		input     *io.Scanner
		err       error
		f         *os.File
		whitelist []int
		key       int
	)
	if len(os.Args) < 2 {
		fmt.Println("args")
		os.Exit(1)
	}
	fname := os.Args[1]
	if f, err = os.Open(fname); err != nil {
		fmt.Println("dud file")
		os.Exit(1)
	}
	scanner = io.NewScanner(f)
	input = io.NewScanner(os.Stdin)
	if whitelist, err = scanner.ReadAllInts(); err != nil {
		fmt.Println("dud whitelist")
		os.Exit(1)
	}
	sort.Ints(whitelist)

	key, err = input.ReadInt()
	for err == nil {
		if rank(key, whitelist) == -1 {
			fmt.Println(key)
		}
		key, err = input.ReadInt()
	}
}

func rank(key int, a []int) int {

	var (
		lo  int
		hi  = len(a) - 1
		mid int
	)

	for lo <= hi {
		mid = lo + (hi-lo)/2
		if key < a[mid] {
			hi = mid - 1
		} else if key > a[mid] {
			lo = mid + 1
		} else {
			return mid
		}
	}
	return -1

}
