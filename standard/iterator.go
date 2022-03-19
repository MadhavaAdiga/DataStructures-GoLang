package standard

type Iterator interface {
	HasNext() bool
	GetNext() interface{}
}
