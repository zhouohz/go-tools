package str

import (
	"fmt"
	"github.com/fogleman/gg"
	"math/rand"
	"testing"
	"time"
)

func TestName(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	const (
		width  = 200
		height = 100
	)

	// 创建一个新的图像上下文
	dc := gg.NewContext(width, height)

	// 设置背景颜色
	dc.SetRGB(1, 1, 1)
	dc.Clear()

	// 设置字体和字体大小
	if err := dc.LoadFontFace("./fonts/Roboto-Medium.ttf", 30); err != nil {
		fmt.Println("Failed to load font:", err)
		return
	}

	// 随机生成验证码文本
	code := generateRandomCode(4)

	// 设置文本颜色
	dc.SetRGB(0, 0, 0)

	//// 计算单个字符的宽度
	//charWidth, _ := dc.MeasureString("A")
	//
	//// 计算验证码文本的总宽度
	//textWidth := charWidth * float64(len(code))
	//
	//// 计算垂直位移的范围
	//maxVerticalOffset := 20
	//
	//// 将绘图原点移到图像中心
	//dc.Translate(width/2, height/2)

	//for i, char := range code {
	//	// 随机垂直位移
	//	verticalOffset := rand.Float64()*float64(maxVerticalOffset*2) - float64(maxVerticalOffset)
	//	// 随机倾斜角度，最多±20度
	//	rotation := rand.Float64()*40 - 20
	//
	//	// 应用垂直位移和倾斜角度
	//	dc.Push()
	//	dc.Rotate(rotation * math.Pi / 180)
	//	dc.Translate(charWidth*float64(i)-textWidth/2, verticalOffset)
	//
	//	// 在当前位置绘制字符
	//	dc.DrawString(string(char), 0, 0)
	//
	//	dc.Pop()
	//}
	dc.DrawStringAnchored(code, width, height, 0.5, 0.5)

	// 画干扰线
	//drawRandomLines(dc, 20)

	// 保存验证码图像
	err := dc.SavePNG("captcha.png")
	if err != nil {
		fmt.Println("Failed to save image:", err)
		return
	}

	fmt.Println("Captcha image saved as captcha.png")
}

func generateRandomCode(length int) string {
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	code := make([]byte, length)
	for i := range code {
		code[i] = charset[rand.Intn(len(charset))]
	}
	return string(code)
}

func drawRandomLines(dc *gg.Context, count int) {
	for i := 0; i < count; i++ {
		x1 := rand.Float64() * float64(dc.Width())
		y1 := rand.Float64() * float64(dc.Height())
		x2 := rand.Float64() * float64(dc.Width())
		y2 := rand.Float64() * float64(dc.Height())
		dc.SetRGB(rand.Float64(), rand.Float64(), rand.Float64())
		dc.DrawLine(x1, y1, x2, y2)
		dc.Stroke()
	}
}
