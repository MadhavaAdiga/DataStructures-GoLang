package bst

import "com.github/MadhavaAdiga/GolangDS/standard/internal/types"

type TreeType int8

type BinarySearchTree interface {
	Add(types.Comparable) bool
	Remove(types.Comparable) bool
	IsEmpty() bool
	Contains(types.Comparable) bool
	Size() int
}
