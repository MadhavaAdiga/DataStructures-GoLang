package linkedlist

type SinglyLinkedList[T any] struct {
	head *node[T]
	tail *node[T]
	size int
}
