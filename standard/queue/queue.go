package queue

import "com.github/MadhavaAdiga/GolangDS/standard/internal/types"

type Queue interface {
	Enqueue(interface{}) error
	Dequeue() (interface{}, error)
	Peek() (interface{}, error)
	elementAt(int) (interface{}, error)
	Size() int
	IsEmpty() bool
	types.Collection
}

type QueueIterator struct {
	index int
	queue Queue
}

func newQueueIterator(queue Queue) *QueueIterator {
	return &QueueIterator{
		index: 0,
		queue: queue,
	}
}

func (iterator *QueueIterator) HasNext() bool {
	return iterator.index < iterator.queue.Size()
}
func (iterator *QueueIterator) GetNext() interface{} {
	v, err := iterator.queue.elementAt(iterator.index)
	if err != nil {
		print(err)
	}
	iterator.index++
	return v
}
