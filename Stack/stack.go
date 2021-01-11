package stack

/*
Stack is interface for stack
*/
type Stack interface {
	Push(element interface{}) error
	Pop(element interface{}) (interface{}, error)
	Size() int
}
