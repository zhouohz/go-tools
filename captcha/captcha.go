package captcha

//type Captcha struct {
//	handler Handle
//}
//
//type Handle interface {
//	BgHandle(image image.Image) any
//	FontHandle(any) any
//}

type Captcha interface {
	// Get 获取验证码
	Get() (GetRes, error)

	// Check 核对验证码
	Check(checkReq CheckReq) (CheckRsp, error)

	//// Verification 二次校验验证码(后端)
	//Verification(validate string) error
}

type GetRes struct {
	Bg    interface{} `json:"bg"`    // 背景图片
	Front interface{} `json:"front"` // 验证
	Token string      `json:"token"` // 唯一标识
	Type  int         `json:"type"`  // 类型
}

type CheckReq struct {
	Token string `json:"token"` // 唯一标识
	Type  int    `json:"type"`
	Data  string `json:"data"` // 数据
}

type CheckRsp struct {
	Token    string `json:"token"` // 唯一标识
	Result   bool   `json:"result"`
	Validate string `json:"validate"`
}
