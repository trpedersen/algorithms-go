package rbtree_string

import (
	"log"
	"strconv"
	"testing"
	"time"
)

const (
	trials int = 1000000
)

func logElapsedTime(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

func TestPutStringKey(t *testing.T) {

	defer logElapsedTime(time.Now(), "TestPutStringKey")

	tree := NewRBTree()

	tree.Put("a", nil)

	value := tree.Get("a")

	if value != nil {
		t.Error(value)
	}
	log.Println(value)

}

func TestPutStringKeys(t *testing.T) {

	tree := NewRBTree()

	f := func() {
		defer logElapsedTime(time.Now(), "TestPutStringKeys")
		for i := 0; i < trials; i++ {

			var s string = strconv.Itoa(i)
			tree.Put(Key(s), nil)

			value := tree.Get(Key(s))

			if value != nil {
				t.Error(Key(s))
			}
		}
		log.Printf("done, size: %d, height: %d\n", tree.Size(), tree.Height())
	}

	for j := 0; j < 10; j++ {
		f()
	}

	//for i

	//log.Printf("tree.Size(): %d\n", tree.Size())
}
