package main

import (
	"bitbucket.org/trpedersen/algorithms/io"
	sort "bitbucket.org/trpedersen/algorithms/sort"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

// func main() {

// 	alg1 := flag.String("a1", "s", "-a1 <algorithm: s - selection, i - insertion>")
// 	alg2 := flag.String("a2", "s", "-a2 <algorithm: s - selection, i - insertion>")
// 	N := flag.Int("N", 1000, "-N <size of array>")
// 	T := flat.Int("T", 100, "-T <number of tests>")

// 	timeSort(*algorithm, a)

// }

func main() {
	useStdIn("mergesort", "string")
}

func main2() {

	//algorithm := flag.String("a", "s", "-a <algorithm: selection, insertion>")
	//datatype := flag.String("dt", "string", "-dt <datatype: string, int")
	//stdIn := flag.Bool("stdin", false, "-stdin")
	alg1 := flag.String("a1", "selection", "-a1 <algorithm: selection, insertion>")
	//alg2 := flag.String("a2", "insertion", "-a2 <algorithm: selection, insertion>")
	//N := flag.Int("N", 1000, "-N <size of array>")
	//T := flag.Int("T", 100, "-T <number of tests>")

	flag.Parse()

	// if *stdIn {
	// 	useStdIn(*algorithm, *datatype)
	// } else {
	// 	compare(*alg1, *alg2, *N, *T)
	// }

	timesA1 := make([]string, 0) //make(map[int]time.Duration)
	//timesA2 := make([]string, 0) //make(map[int]time.Duration)
	for N := 1000; N <= 100000; N = N + 1000 {
		//compare(*alg1, *alg2, N, *T)
		timesA1 = append(timesA1, fmt.Sprintf("%d:%s, ", N, timeRandom(*alg1, N, 1)))
		//	timesA2 = append(timesA2, fmt.Sprintf("%d:%s, ", N, timeRandom(*alg2, N, 1)))
	}
	log.Println(timesA1)
	//	log.Println(timesA2)
}

func compare(alg1 string, alg2 string, N int, T int) {
	d1 := timeRandom(alg1, N, T)
	d2 := timeRandom(alg2, N, T)

	log.Printf("N: %d, T: %d", N, T)
	log.Printf("alg1: %s, duration: %s", alg1, d1)
	log.Printf("alg2: %s, duration: %s", alg2, d2)
}

func useStdIn(algorithm string, datatype string) {

	stdIn := io.NewScanner(os.Stdin)

	//var result sort.Interface
	var a sort.Interface

	switch datatype {
	case "string":
		log.Print("Reading tokens...")
		tokens, err := stdIn.ReadAllTokens()
		if err != nil {
			panic(err)
		}
		log.Printf("done, tokens: %d", len(tokens))
		// for _, token := range tokens {
		// 	fmt.Println(token)
		// }
		a = sort.Strings(tokens)
	case "int":
		log.Print("Reading ints...")
		ints, err := stdIn.ReadAllInts()
		if err != nil {
			panic(err)
		}
		log.Printf("done, ints: %d", len(ints))
		a = sort.Ints(ints)
	}
	log.Println("Starting sort")
	t := timeSort(algorithm, a)

	//log.Println(a)
	log.Printf("Sort done, algorithm: %s time: %s\n", algorithm, t)
	//os.Exit(result.Len())
	//fmt.Println(result)

}

func makeRandomIntArray(N int) sort.Ints {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	a := make([]int, N, N)
	for i := 0; i < N; i++ {
		//a = append(a, r.Intn(9999))
		a[i] = r.Intn(N)
	}
	return a
}

func timeRandom(algorithm string, N int, T int) time.Duration {

	var total time.Duration
	for t := 0; t < T; t++ {
		a := makeRandomIntArray(N)
		total += timeSort(algorithm, a)
	}
	return time.Duration(total.Nanoseconds() / int64(T))
}

func timeSort(algorithm string, a sort.Interface) time.Duration {
	var f func(sort.Interface) sort.Interface

	switch algorithm {
	case "selection":
		f = sort.SelectionSort
	case "insertion":
		f = sort.InsertionSort
	case "shellsort":
		f = sort.ShellSort
	case "mergesort":
		f = sort.MergeSort
	}
	start := time.Now()
	f(a)
	elapsed := time.Since(start)
	return elapsed
}

func timer(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}
