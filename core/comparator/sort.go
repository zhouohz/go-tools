package comparator

type Sort interface {
	Compare(a, b any) int
}
