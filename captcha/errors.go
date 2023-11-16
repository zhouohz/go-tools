package captcha

import "errors"

var (
	InvalidCodeError = errors.New("无效验证码")
	CheckCodeError   = errors.New("验证码错误")
)
