package library

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {

	fmt.Println(GetPinyinEngine().Sentence("魃魈魁鬾魑魅魍魎").Unicode())
}
