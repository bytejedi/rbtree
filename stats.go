package rbtree

// This file contains most of the methods that can be used
// by the user. Anyone who wants to look for some API about
// the rbtree, this is the right place.

// Number of nodes in the tree.
func (t *Rbtree) Len() uint {
	t.mu.RLock()
	defer t.mu.RUnlock()
	return t.count
}

func (t *Rbtree) Insert(item Item) {
	t.mu.Lock()
	defer t.mu.Unlock()
	if item == nil {
		return
	}

	// Always insert a RED node
	t.insert(&Node{t.NIL, t.NIL, t.NIL, RED, item})
}

//InsertOrGet inserts or retrieves the item in the tree. If the
//item is already in the tree then the return value will be that.
//If the item is not in the tree the return value will be the item
//you put in.
func (t *Rbtree) InsertOrGet(item Item) Item {
	t.mu.Lock()
	defer t.mu.Unlock()
	if item == nil {
		return nil
	}

	return t.insert(&Node{t.NIL, t.NIL, t.NIL, RED, item}).Item
}

func (t *Rbtree) Delete(item Item) Item {
	t.mu.Lock()
	defer t.mu.Unlock()
	if item == nil {
		return nil
	}

	// The `color` field here is nobody
	return t.delete(&Node{t.NIL, t.NIL, t.NIL, RED, item}).Item
}

func (t *Rbtree) Get(item Item) Item {
	t.mu.RLock()
	defer t.mu.RUnlock()
	if item == nil {
		return nil
	}

	// The `color` field here is nobody
	ret := t.search(&Node{t.NIL, t.NIL, t.NIL, RED, item})
	if ret == nil {
		return nil
	}

	return ret.Item
}

func (t *Rbtree) Min() Item {
	t.mu.RLock()
	defer t.mu.RUnlock()
	x := t.min(t.root)

	if x == t.NIL {
		return nil
	}

	return x.Item
}

func (t *Rbtree) Max() Item {
	t.mu.RLock()
	defer t.mu.RUnlock()
	x := t.max(t.root)

	if x == t.NIL {
		return nil
	}

	return x.Item
}
