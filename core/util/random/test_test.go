package random

import (
	"testing"
)

func TestName(t *testing.T) {
	RandNumberInMainInterval(struct{ start, end int }{start: 10, end: 20}, []struct{ start, end int }{{12, 18}})
}
