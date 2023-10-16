package dfa

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	filter := New()

	filter.AddWord("垃圾")

	replace := filter.Replace("田鹏真垃圾", '*')

	fmt.Println(replace)
}
