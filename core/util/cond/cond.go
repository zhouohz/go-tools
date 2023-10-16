package cond

func If[T any](condition bool, trueVal, falseVal T) T {
	if condition {
		return trueVal
	}
	return falseVal
}

func IfF[T any](condition bool, ifFunc func() T, elseFunc func() T) T {
	if condition {
		return ifFunc()
	}

	return elseFunc()
}
