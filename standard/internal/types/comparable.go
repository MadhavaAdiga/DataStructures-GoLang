package types

type Comparer func(a Comparable) int32

type Comparable struct {
	Val     interface{}
	Compare Comparer
}

func NewComparable(v interface{}, comparer Comparer) *Comparable {
	return &Comparable{Val: v, Compare: comparer}
}
