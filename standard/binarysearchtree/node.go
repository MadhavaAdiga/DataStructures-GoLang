package binarysearchtree

import "com.github/MadhavaAdiga/GolangDS/standard"

// type Comparer func(c Comparable) int32

// type Comparable struct {
// 	Val     interface{}
// 	Compare Comparer
// }

// func NewComparable(v interface{}, comparer standard.Comparer) * standard.Comparable {
// 	return &standard.Comparable{Val: v, Compare: comparer}
// }

type primitiveNode struct {
	data  standard.Comparable
	left  *primitiveNode
	right *primitiveNode
}

func newPNode(data standard.Comparable) *primitiveNode {
	return &primitiveNode{data, nil, nil}
}
