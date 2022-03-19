package standard

type Collection interface {
	CreateIterator() Iterator
}
