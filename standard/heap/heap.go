package heap

type Heap interface {
	Add(interface{}) error
	Poll() (interface{}, error)
	Peek() (interface{}, error)
	Remove(interface{}) (interface{}, error)
	removeAt(int) (interface{}, error)
	Size() int
	IsEmpty() bool
	Contains(interface{}) bool
	Clear()
}
