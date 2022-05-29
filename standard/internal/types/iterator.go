package types

type Iterator interface {
	HasNext() bool
	GetNext() interface{}
}
