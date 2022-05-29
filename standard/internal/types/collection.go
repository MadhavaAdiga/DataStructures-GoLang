package types

type Collection interface {
	CreateIterator() Iterator
}
