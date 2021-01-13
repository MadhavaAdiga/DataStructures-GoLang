package queue

/*
Queue is an interface to implement queue
*/
type Queue interface {
	Enqueue(value interface{}) error
	Dequeue() (interface{}, error)
	Poll() (interface{}, error)
	Size() int
}
