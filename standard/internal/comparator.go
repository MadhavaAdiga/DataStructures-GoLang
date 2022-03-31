package internal

type Comparator func(a, b interface{}) int

func GetComparator(a interface{}) Comparator {
	switch a.(type) {
	case int:
		return IntComparator
	case int8:
		return Int8Comparator
	case int16:
		return Int16Comparator
	case int32:
		return Int32Comparator
	case int64:
		return Int64Comparator

	case uint:
		return UintComparator
	case uint8:
		return Uint8Comparator
	case uint16:
		return Uint16Comparator
	case uint32:
		return Uint32Comparator
	case uint64:
		return Uint64Comparator

	case float32:
		return Float32Comparator
	case float64:
		return Float64Comparator

	default:
		return StringComparator
	}
}

func IntComparator(a, b interface{}) int {
	v1 := a.(int)
	v2 := b.(int)

	switch {
	case v1 < v2:
		return -1
	case v1 > v2:
		return 1
	default:
		return 0
	}
}

func Int8Comparator(a, b interface{}) int {
	v1 := a.(int8)
	v2 := b.(int8)

	switch {
	case v1 < v2:
		return -1
	case v1 > v2:
		return 1
	default:
		return 0
	}
}

func Int16Comparator(a, b interface{}) int {
	v1 := a.(int16)
	v2 := b.(int16)

	switch {
	case v1 < v2:
		return -1
	case v1 > v2:
		return 1
	default:
		return 0
	}
}

func Int32Comparator(a, b interface{}) int {
	v1 := a.(int32)
	v2 := b.(int32)

	switch {
	case v1 < v2:
		return -1
	case v1 > v2:
		return 1
	default:
		return 0
	}
}

func Int64Comparator(a, b interface{}) int {
	v1 := a.(int64)
	v2 := b.(int64)

	switch {
	case v1 < v2:
		return -1
	case v1 > v2:
		return 1
	default:
		return 0
	}
}

func Uint8Comparator(a, b interface{}) int {
	v1 := a.(uint8)
	v2 := b.(uint8)

	switch {
	case v1 < v2:
		return -1
	case v1 > v2:
		return 1
	default:
		return 0
	}
}

func Uint16Comparator(a, b interface{}) int {
	v1 := a.(uint16)
	v2 := b.(uint16)

	switch {
	case v1 < v2:
		return -1
	case v1 > v2:
		return 1
	default:
		return 0
	}
}

func Uint32Comparator(a, b interface{}) int {
	v1 := a.(uint32)
	v2 := b.(uint32)

	switch {
	case v1 < v2:
		return -1
	case v1 > v2:
		return 1
	default:
		return 0
	}
}

func Uint64Comparator(a, b interface{}) int {
	v1 := a.(uint64)
	v2 := b.(uint64)

	switch {
	case v1 < v2:
		return -1
	case v1 > v2:
		return 1
	default:
		return 0
	}
}

func UintComparator(a, b interface{}) int {
	v1 := a.(uint)
	v2 := b.(uint)

	switch {
	case v1 < v2:
		return -1
	case v1 > v2:
		return 1
	default:
		return 0
	}
}

func Float32Comparator(a, b interface{}) int {
	v1 := a.(float32)
	v2 := b.(float32)

	switch {
	case v1 < v2:
		return -1
	case v1 > v2:
		return 1
	default:
		return 0
	}
}

func Float64Comparator(a, b interface{}) int {
	v1 := a.(float64)
	v2 := b.(float64)

	switch {
	case v1 < v2:
		return -1
	case v1 > v2:
		return 1
	default:
		return 0
	}
}

func StringComparator(a, b interface{}) int {
	v1 := a.(string)
	v2 := b.(string)

	// min := len(v2)
	// if len(v1) < len(v2) {
	// 	min = len(v1)
	// }
	// diff := 0
	// for i := 0; i < min && diff == 0; i++ {
	// 	diff = int(v1[i]) - int(v2[i])
	// }
	// if diff == 0 {
	// 	diff = len(v1) - len(v2)
	// }
	// if diff < 0 {
	// 	return -1
	// }
	// if diff > 0 {
	// 	return 1
	// }
	// return 0

	switch {
	case v1 < v2:
		return -1
	case v1 > v2:
		return 1
	default:
		return 0
	}
}

func StructComparator(a, b interface{}) int {
	v1 := a.(float64)
	v2 := b.(float64)

	switch {
	case v1 < v2:
		return -1
	case v1 > v2:
		return 1
	default:
		return 0
	}
}
