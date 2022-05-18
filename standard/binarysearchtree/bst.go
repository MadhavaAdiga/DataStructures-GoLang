package binarysearchtree

type TreeType int8

const (
	undefined TreeType = iota
	PRIMITIVE
	CUSTOME
)

type BinarySearchTree interface {
	Add(Comparable) bool
	Remove(Comparable) bool
	IsEmpty() bool
	Contains(Comparable) bool
	Size() int
}

// func GetBst(t TreeType) BinarySearchTree {

// 	// switch t {
// 	// case PRIMITIVE:
// 	// 	return newPrimitiveImpl()

// 	// case CUSTOME:
// 	// 	return newCustomImpl()

// 	// default:
// 	// 	return nil
// 	// }
// }
