package rbtree_string

import ()

type Key string

type Value *interface{}

type Node struct {
	key    Key
	value  Value
	left   *Node
	right  *Node
	colour bool
	N      int
}

const (
	RED   bool = true
	BLACK bool = false
)

func NewNode(key Key, value Value, colour bool, N int) *Node {
	return &Node{
		key:    key,
		value:  value,
		colour: colour,
		N:      N,
	}
}

func isRed(node *Node) bool {
	if node == nil {
		return false
	} else {
		return node.colour == RED
	}
}

type RBTree struct {
	root     *Node
	compares int
}

func NewRBTree() *RBTree {
	return &RBTree{}
}

func (tree *RBTree) Size() int {
	return size(tree.root)
}

func (tree *RBTree) Height() int {
	if tree.root == nil {
		return 0
	} else {
		return height(tree.root)
	}
}

func min(a int, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}

func max(a int, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

func height(node *Node) int {
	if node == nil {
		return -1
	}
	return 1 + max(height(node.left), height(node.right))
}

func (tree *RBTree) Put(key Key, value Value) {
	tree.root = put(tree.root, key, value)
}

func compare(a Key, b Key) int {
	//compares += 1
	if a == b {
		return 0
	} else if a < b {
		return -1
	} else {
		return 1
	}
}

func size(node *Node) int {
	if node == nil {
		return 0
	}
	return node.N
}

func put(node *Node, key Key, value Value) *Node {
	// change key's value to value if key in subtree rooted at node
	// otherwise add a new node to subtree associating key with value.
	if node == nil {
		return NewNode(key, value, RED, 1)
	}
	cmp := compare(key, node.key)
	if cmp < 0 {
		node.left = put(node.left, key, value)
	} else if cmp > 0 {
		node.right = put(node.right, key, value)
	} else {
		node.key = key
	}

	if isRed(node.right) && !isRed(node.left) {
		node = rotateLeft(node)
	}
	if isRed(node.left) && isRed(node.left.left) {
		node = rotateRight(node)
	}
	if isRed(node.left) && isRed(node.right) {
		flipColours(node)
	}
	node.N = 1 + size(node.left) + size(node.right)

	return node
}

func rotateLeft(h *Node) *Node {
	x := h.right
	h.right = x.left
	x.left = h
	x.colour = h.colour
	h.colour = RED
	x.N = h.N
	h.N = 1 + size(h.left) + size(h.right)
	return x
}

func rotateRight(h *Node) *Node {
	x := h.left
	h.left = x.right
	x.right = h
	x.colour = h.colour
	h.colour = RED
	x.N = h.N
	h.N = 1 + size(h.left) + size(h.right)
	// rotateRights++
	return x
}

func flipColours(h *Node) {
	h.colour = !h.colour
	h.left.colour = !h.left.colour
	h.right.colour = !h.right.colour
	//flipColours++
}

func (tree *RBTree) Get(key Key) Value {
	node := get(tree.root, key)
	if node != nil {
		return node.value
	} else {
		return nil
	}
}

func get(node *Node, key Key) *Node {
	if node == nil {
		return nil
	}
	cmp := compare(key, node.key)
	if cmp < 0 {
		return get(node.left, key)
	} else if cmp > 0 {
		return get(node.right, key)
	} else {
		return node
	}
}
