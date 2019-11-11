package main

import (
	"fmt"
	"github.com/bytejedi/rbtree"
)

func main() {
	rbt := rbtree.New()

	rbt.Insert(rbtree.String("Hello"))
	rbt.Insert(rbtree.String("World"))

	rbt.Ascend(rbt.Min(), Print)
}

func Print(item rbtree.Item) bool {
	i, ok := item.(rbtree.String)
	if !ok {
		return false
	}
	fmt.Println(i)
	return true
}
