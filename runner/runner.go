package main

import (
	"bitbucket.org/trpedersen/algorithms/io"
	"bitbucket.org/trpedersen/algorithms/stats"
	"fmt"
	"os"
	"sort"
)

func main() {
	//	f, err := os.Open("C:\\Users\\timpe_000\\SkyDrive\\go\\data\\largeW.txt")
	//f, err := os.Open("/Users/tpedersen/SkyDrive/go/data/largeW.txt")
	//	if(err != nil){
	//		os.Exit(1)
	//	}
	stdIn := io.NewScanner(os.Stdin)

	ints, _ := stdIn.ReadAllInts()

	sort.Ints(ints)

	var x int64
	var accum = stats.NewInt64Accumulator()

	for i := range ints {
		//fmt.Println(i)
		x = x + int64(i)
		accum.Add(int64(i))
	}
	fmt.Println(x)
	fmt.Println(accum)
	/*
		i, err := stdIn.ReadString()
		for err == nil {
			fmt.Println(i)
			if i, err = stdIn.ReadString(); err != nil {
				fmt.Println(i, err)
			}
		}
	*/
}
