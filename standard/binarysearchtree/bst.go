package binarysearchtree

import "com.github/MadhavaAdiga/GolangDS/standard"

type TreeType int8

const (
	undefined TreeType = iota
	PRIMITIVE
	CUSTOME
)

type BinarySearchTree interface {
	Add(standard.Comparable) bool
	Remove(standard.Comparable) bool
	IsEmpty() bool
	Contains(standard.Comparable) bool
	Size() int
}
