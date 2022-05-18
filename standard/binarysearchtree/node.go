package binarysearchtree

// type Node interface {
// 	isEmpty() bool
// }

type Comparer func(c Comparable) int32

type Comparable struct {
	Val     interface{}
	Compare Comparer
}

func NewComparable(v interface{}, comparer Comparer) *Comparable {
	return &Comparable{Val: v, Compare: comparer}
}

type primitiveNode struct {
	data  Comparable
	left  *primitiveNode
	right *primitiveNode
}

func newPNode(data Comparable) *primitiveNode {
	return &primitiveNode{data, nil, nil}
}
