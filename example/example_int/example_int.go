package main

import (
	"fmt"
	"github.com/bytejedi/rbtree"
)

func main() {
	rbt := rbtree.New()

	m := 0
	n := 10

	for m < n {
		rbt.Insert(rbtree.Int(m))
		m++
	}

	m = 0
	for m < n {
		if m%2 == 0 {
			rbt.Delete(rbtree.Int(m))
		}
		m++
	}

	rbt.Ascend(rbt.Min(), Print)
}

func Print(item rbtree.Item) bool {
	i, ok := item.(rbtree.Int)
	if !ok {
		return false
	}
	fmt.Println(i)
	return true
}
