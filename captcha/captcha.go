package captcha

import (
	"context"
)

const (
	Click = iota + 1
	Puzzle
	Seq
	Slide
)

const (
	answerSuffix = "captcha_"
)

func GetID(s string) string {
	return answerSuffix + s
}

type Captcha interface {
	// Get 获取验证码
	Get(ctx context.Context) (*Generate, error)

	// Check 核对验证码
	Check(ctx context.Context, token, pointJson string) error

	//// Verification 二次校验验证码(后端)
	//Verification(validate string) error
}

type Generate struct {
	Bg     interface{} `json:"bg"`     // 背景图片
	Front  interface{} `json:"front"`  // 验证
	Token  string      `json:"token"`  // 唯一标识
	Answer interface{} `json:"answer"` // 答案
}
