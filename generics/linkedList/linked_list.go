package linkedlist

type LinkedList[T any] interface {
	Add(T) error
	AddFirst(T) error
	Remove() (T, error)
	RemoveAt(int) (T, error)
	RemoveFirst() (T, error)
	IndexOf(T) (int, error)
	Contains(T) bool
	Size() int
	IsEmpty() bool
}
