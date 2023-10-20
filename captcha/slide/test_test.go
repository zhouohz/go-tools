package slide

import (
	"github.com/fogleman/gg"
	"image"
	"image/color"
	"image/png"
	"os"
	"testing"
)

func TestName(t *testing.T) {
	// 1. 打开图像文件
	imageFile, err := os.Open("D:\\work\\go-tools\\captcha\\slide\\res\\bg\\1.png")
	if err != nil {
		panic(err)
	}
	defer imageFile.Close()

	// 2. 解码图像文件
	img, _, err := image.Decode(imageFile)
	if err != nil {
		panic(err)
	}

	// 3. 创建图像上下文
	context := gg.NewContextForImage(img)

	// 4. 创建蒙版并设置剪切区域
	mask := createMask(context, 50, 60, 100, 100) // 指定剪切区域

	// 5. 将蒙版应用到图像上下文
	if err := context.SetMask(mask); err != nil {
		panic(err)
	}

	// 6. 执行剪切操作
	context.Clip()

	// 7. 保存裁剪后的图像
	outputImage, err := os.Create("output.png")
	if err != nil {
		panic(err)
	}
	defer outputImage.Close()

	if err := png.Encode(outputImage, context.Image()); err != nil {
		panic(err)
	}

	// 8. 保存裁剪区域为新图像
	maskImage, err := os.Create("mask.png")
	if err != nil {
		panic(err)
	}
	defer maskImage.Close()

	if err := png.Encode(maskImage, mask); err != nil {
		panic(err)
	}
}

func createMask(context *gg.Context, x1, y1, x2, y2 int) *image.Alpha {
	// 创建一个蒙版，以指定区域为不透明，其余区域为透明
	bounds := context.Image().Bounds()
	mask := image.NewAlpha(bounds)

	for x := x1; x < x2; x++ {
		for y := y1; y < y2; y++ {
			mask.SetAlpha(x, y, color.Alpha{255})
		}
	}

	return mask
}
