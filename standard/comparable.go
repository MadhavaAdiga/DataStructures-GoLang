package standard

type Comparer func(a interface{}) int32

type Comparable struct {
	val     interface{}
	compare Comparer
}

func NewComparable(comparer Comparer) *Comparable {
	return &Comparable{compare: comparer}
}
