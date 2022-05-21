package binarysearchtree

import (
	"reflect"

	"com.github/MadhavaAdiga/GolangDS/standard"
)

type BstImpl struct {
	nodeCount int
	root      *primitiveNode
	// comparator internal.Comparator
}

func NewPrimitiveImpl() *BstImpl {
	return &BstImpl{}
}

func (tree *BstImpl) Add(ele standard.Comparable) bool {
	if !tree.typeCheck(ele) {
		return false
	}

	// handle duplicate
	if tree.Contains(ele) {
		return false
	}

	var prev *primitiveNode = nil // parent pointer of traversing node
	curr := tree.root             // traversing node

	// find the leafnode to insert
	for curr != nil {
		prev = curr
		// cmp := tree.comparator(ele, curr.data)
		cmp := ele.Compare(curr.data)

		if cmp < 0 {
			curr = curr.left
		} else if cmp > 0 {
			curr = curr.right
		}
	}

	node := newPNode(ele)

	if prev != nil {
		// cmp := tree.comparator(ele, prev.data)
		cmp := ele.Compare(prev.data)

		if cmp < 0 {
			prev.left = node
		} else if cmp > 0 {
			prev.right = node
		}
	} else {
		tree.root = node
		// tree.comparator = internal.GetComparator(ele)
	}
	tree.nodeCount++
	return true
}

func (tree *BstImpl) Remove(ele standard.Comparable) bool {
	if !tree.typeCheck(ele) {
		return false
	}

	curr := tree.root             // node to be found
	var prev *primitiveNode = nil // parent of found node

	// find the node to delete
	for curr != nil {
		// cmp := tree.comparator(ele, curr.data)
		cmp := ele.Compare(curr.data)

		if cmp < 0 {
			prev = curr
			curr = curr.left
		} else if cmp > 0 {
			prev = curr
			curr = curr.right
		} else if cmp == 0 {
			break
		}
	}

	if curr == nil {
		return false
	}

	// node is found
	// handle diffrent cases for removal
	if curr.left != nil && curr.right != nil {
		// both subtree exists
		// choose either one

		//find smallest in right sub tree
		var parent *primitiveNode = nil // parent of the selected child
		pointer := curr.right           // found nodes right child, used to find the leaf node in selected subtree

		// find the smallest leaf node
		for pointer.left != nil {
			parent = pointer
			pointer = pointer.left
		}

		// check if parent is not the current node ( parent == nil case occures when found node's subtree has height of 1 )
		// replace the (leafnode) parent's left with nil
		if parent != nil {
			// parent.left = pointer.right
			parent.left = nil
		} else {
			// when parent == nil ( case occures when found node's subtree has height of 1 )
			// replace found node's right with found node's right's right
			curr.right = pointer.right
		}

		// fin the largest leaf node
		// var parent *primitiveNode = nil // parent of the selected child
		// pointer := curr.left           // found nodes right child, used to find the leaf node in selected subtree

		// for pointer.right != nil {
		// 	parent = pointer
		// 	pointer = pointer.right
		// }
		// if parent != nil {
		// 	parent.right = nil
		// } else {{
		// 	curr.left = pointer.left
		// }

		// copy the leafnode data to found node
		// removal of found node is done by replacing data with either the
		// smallest child in right subtree
		// or
		// largest child in leftsubtree
		curr.data = pointer.data
	} else if curr.left != nil || curr.right != nil {
		var pointer *primitiveNode

		// case where either only left or right sub tree exists
		if curr.left == nil {
			pointer = curr.right
		} else if curr.right == nil {
			pointer = curr.left
		}

		// check if current node is previous node's left or right child
		// replace the current node with pointer node
		if curr == prev.left {
			prev.left = pointer
		} else {
			prev.right = pointer
		}
	} else if curr.left == nil && curr.right == nil {
		// when node to be removed is the leaf node
		if prev == nil {
			tree.root = nil
		} else {
			if curr == prev.left { // check if found node(curr) is parents left or right child
				prev.left = nil
			} else {
				prev.right = nil
			}
		}
	}
	tree.nodeCount--

	return true
}

func (tree *BstImpl) IsEmpty() bool {
	// return (tree.root == primitiveNode{})
	return tree.root == nil
}

func (tree *BstImpl) Contains(ele standard.Comparable) bool {
	pointer := tree.root

	for pointer != nil {
		// cmp := tree.comparator(ele, pointer.data)
		cmp := ele.Compare(pointer.data)

		if cmp < 0 {
			pointer = pointer.left
		} else if cmp > 0 {
			pointer = pointer.right
		} else {
			return true
		}
	}

	return false
}

func (tree *BstImpl) Size() int {
	return tree.nodeCount
}

func (tree *BstImpl) typeCheck(ele standard.Comparable) bool {
	if tree.IsEmpty() {
		return true
	}
	v := tree.root.data
	return reflect.TypeOf(ele) == reflect.TypeOf(v)
}
