package binarysearchtree

import "com.github/MadhavaAdiga/GolangDS/standard"

// type Node interface {
// 	isEmpty() bool
// }

type primitiveNode struct {
	data  interface{}
	left  *primitiveNode
	right *primitiveNode
}

func newPNode(data interface{}) *primitiveNode {
	return &primitiveNode{data, nil, nil}
}

// func (node *primitiveNode) isEmpty() bool {
// 	return node.data == nil && node.left == nil && node.right == nil
// }

type customeNode struct {
	data  standard.Comparable
	left  *customeNode
	right *customeNode
}

func newCNode(data standard.Comparable) *customeNode {
	return &customeNode{data, nil, nil}
}

// var _ Node = (*primitiveNode)(nil)
