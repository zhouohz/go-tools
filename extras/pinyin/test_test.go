package library

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {

	fmt.Println(GetPinyinEngine().Sentence("倚天屠龙记").Unicode())
}
