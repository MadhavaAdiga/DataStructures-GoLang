package binarysearchtree

type TreeType int8

const (
	undefined TreeType = iota
	PRIMITIVE
	CUSTOME
)

type BinarySearchTree interface {
	Add(interface{}) bool
	Remove(interface{}) bool
	IsEmpty() bool
	Contains(interface{}) bool
	Size() int
}

func GetBst(t TreeType) BinarySearchTree {

	switch t {
	case PRIMITIVE:
		return newPrimitiveImpl()

	case CUSTOME:
		return newCustomImpl()

	default:
		return nil
	}
}
