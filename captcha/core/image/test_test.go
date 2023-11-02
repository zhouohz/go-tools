package image

import (
	"fmt"
	"os"
	"testing"
)

func TestName(t *testing.T) {
	character := NewCharacter(100, 256, 1, 1, 4, "", nil)

	id, content, answer := character.GenerateIdQuestionAnswer()

	fmt.Println(id)
	fmt.Println(content)
	fmt.Println(answer)

	item, err := character.DrawCaptcha(content)
	if err != nil {
		fmt.Println(err)
		return
	}

	file, err := os.Create("./captcha.png")
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	item.WriteTo(file)
}
