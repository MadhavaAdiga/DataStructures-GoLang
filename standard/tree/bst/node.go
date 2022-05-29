package bst

import "com.github/MadhavaAdiga/GolangDS/standard/internal/types"

type node struct {
	data  types.Comparable
	left  *node
	right *node
}

func newPNode(data types.Comparable) *node {
	return &node{data, nil, nil}
}
