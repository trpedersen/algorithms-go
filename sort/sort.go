package sort

import (
	//"fmt"
	"log"
)

type Interface interface {

	// Len is the number of elements in the collection.
	Len() int
	// Less reports whether the element with
	// index i should sort before the element with index j.
	Less(i, j int) bool
	// Swap swaps the elements with indexes i and j.
	Swap(i, j int)
	// Clone
	Copy(b []interface{}) // Interface

	Get(i int) interface{}
	Set(i int, item interface{})
}

type comparator func(a []interface{}, i int, j int) bool

type Strings []string

func (a Strings) Len() int           { return len(a) }
func (a Strings) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Strings) Less(i, j int) bool { return a[i] < a[j] }
func (a Strings) Copy(b []interface{}) {
	//copy(b, a)
	for i, v := range a {
		b[i] = v
	}
}

// func (a Strings) Copy(b []interface{})  {
// 	//clone := make([]string, len(a))
// 	copy(b, a)
// 	return Strings(b)
// }

func (a Strings) Get(i int) interface{} {
	return a[i]
}
func (a Strings) Set(i int, value interface{}) {
	a[i] = value.(string)
}

type Ints []int

func (a Ints) Len() int           { return len(a) }
func (a Ints) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Ints) Less(i, j int) bool { return a[i] < a[j] }
func (a Ints) Copy(b []interface{}) {
	//copy(b, a)
	for i, v := range a {
		b[i] = v
	}
}

// func (a Ints) Copy() Interface {
// 	clone := make([]int, len(a))
// 	copy(clone, a)
// 	return Ints(clone)
// }
func (a Ints) Get(i int) interface{} {
	return a[i]
}
func (a Ints) Set(i int, value interface{}) {
	a[i] = value.(int)
}

func less(a interface{}, b interface{}) bool {
	switch a.(type) {
	case string:
		return a.(string) < b.(string)
	case int:
		return a.(int) < b.(int)
	default:
		return true
	}
}

func IsSortedAsc(a Interface) bool {
	for i := 1; i < a.Len(); i++ {
		if a.Less(i, i-1) {
			return false
		}
	}
	return true
}

func IsSortedDesc(a Interface) bool {
	for i := a.Len() - 1; i > 0; i-- {
		if a.Less(i-1, i) {
			return false
		}
	}
	return true
}

// SelectionSort sorts a[] in ascending order
func SelectionSort(a Interface) Interface {
	N := a.Len()
	for i := 0; i < N; i++ {
		// exchange a[i] with smallest entry in a[i+1:N]
		min := i
		for j := i + 1; j < N; j++ {
			if a.Less(j, min) {
				min = j
			}
		}
		a.Swap(i, min)
	}
	return a
}

// InsertionSort sorts a[] in ascending order
func InsertionSort(a Interface) Interface {
	N := a.Len()
	for i := 0; i < N; i++ {
		for j := i; j > 0 && a.Less(j, j-1); j-- {
			a.Swap(j, j-1)
		}
	}
	return a
}

func ShellSort(a Interface) Interface {
	N := a.Len()
	h := 1

	swaps := 0

	for h < N/3 {
		h = 3*h + 1 // 1, 4, 13, 40, 121, 364, 1093, ...
	}
	for h >= 1 {
		// h-sort the array
		for i := h; i < N; i++ {
			// insert a[i] among a[i-h], a[1-2*h], a[i-3*h], ...
			for j := i; j >= h && a.Less(j, j-h); j -= h {
				a.Swap(j, j-h)
				swaps++
				if swaps%100000 == 0 {
					log.Printf("swaps: %d", swaps)
				}
			}
		}
		h = h / 3
	}
	return a
}

var aux []interface{}

// Merge a[lo...mid] with a[mid+1..hi]
func merge(a Interface, lo int, mid int, hi int) Interface {
	i := lo
	j := mid + 1
	//copy(aux, a) // = a.Clone()
	//a.Copy(aux)
	for k := lo; k <= hi; k++ {
		aux[k] = a.Get(k)
		//fmt.Print(aux[k])
	}

	//fmt.Println()

	for k := lo; k <= hi; k++ {
		if i > mid {
			// lhs exhausted, take from rhs
			a.Set(k, aux[j])
			j++
		} else if j > hi {
			// rhs exhausted, take from lhs
			a.Set(k, aux[i])
			i++
		} else if less(aux[j], aux[i]) {
			// take from rhs
			a.Set(k, aux[j])
			j++
		} else {
			// take from lhs
			a.Set(k, aux[i])
			i++
		}
	}
	//fmt.Println(a)
	return a
}

// // Merge a[lo...mid] with a[mid+1..hi]
// func merge(a Interface, lo int, mid int, hi int) Interface {
// 	i := lo
// 	j := mid + 1
// 	//copy(aux, a) // = a.Clone()
// 	aux = a.Clone()

// 	for k := lo; k <= hi; k++ {
// 		if i > mid {
// 			// lhs exhausted, take from rhs
// 			a[k] = aux[j]
// 			j++
// 		} else if j > hi {
// 			// rhs exhausted, take from lhs
// 			a[k] = aux[i]
// 			i++
// 		} else if a.Less(j, i) {
// 			// take from rhs
// 			a[k] = aux[j]
// 			j++
// 		} else {
// 			// take from lhs
// 			a[k] = aux[i]
// 			i++
// 		}
// 	}
// 	return a
// }

// Sort a[lo..hi]
func MergeSort(a Interface) Interface {
	aux = make([]interface{}, a.Len())
	mergeSort(a, 0, a.Len()-1)
	//fmt.Println(a)
	return a
}

func mergeSort(a Interface, lo int, hi int) Interface {
	if hi <= lo {
		return a
	}
	mid := lo + (hi-lo)/2
	mergeSort(a, lo, mid)
	mergeSort(a, mid+1, hi)
	//fmt.Println("1", a, lo, mid, hi)
	merge(a, lo, mid, hi)
	//fmt.Println("2", a)
	return a
}
