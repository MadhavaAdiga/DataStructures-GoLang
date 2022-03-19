package linkedlist

type node[T any] struct {
	data T
	prev *node[T]
	next *node[T]
}

func newNode[T any](data T) *node[T] {
	return &node[T]{
		data: data,
		next: nil,
		prev: nil,
	}
}
