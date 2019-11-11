package rbtree

type Int int

func (x Int) Less(than Item) bool {
	return x < than.(Int)
}

type String string

func (x String) Less(than Item) bool {
	return x < than.(String)
}
