package image

import (
	"bytes"
	"encoding/base64"
	"golang.org/x/image/draw"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// ParseImage 解析图片并返回解析后的图像对象
// imagePath：图片文件路径
func ParseImage(imagePath string) (image.Image, error) {
	// 打开图片文件
	file, err := os.Open(imagePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var img image.Image
	// 使用文件扩展名来确定图像类型
	if isJPEG(imagePath) {
		img, err = jpeg.Decode(file)
	} else if isPNG(imagePath) {
		img, err = png.Decode(file)
	}

	if err != nil {
		return nil, err
	}

	return img, nil
}

// isJPEG 检查文件扩展名是否为JPEG
func isJPEG(filename string) bool {
	ext := filepath.Ext(filename)
	return strings.EqualFold(ext, ".jpg") || strings.EqualFold(ext, ".jpeg")
}

// isPNG 检查文件扩展名是否为PNG
func isPNG(filename string) bool {
	ext := filepath.Ext(filename)
	return strings.EqualFold(ext, ".png")
}

// CutImage 通过模板图片抠图（不透明部分）
// oriImage：源图片
// templateImage：模板图片
// xPos：坐标轴x
// yPos：坐标轴y
// 返回值：裁剪后的图像
func CutImage(oriImage image.Image, templateImage image.Image, xPos, yPos int) image.Image {
	templateBounds := templateImage.Bounds()
	targetImage := image.NewRGBA(templateBounds)

	for y := 0; y < templateBounds.Dy(); y++ {
		for x := 0; x < templateBounds.Dx(); x++ {
			alpha := templateImage.At(x, y).(color.NRGBA).A
			if alpha > 100 {
				bgRgb := oriImage.At(xPos+x, yPos+y)
				targetImage.Set(x, y, bgRgb)
			}
		}
	}

	return targetImage
}

// CutImageWithBorderBlur 通过模板图片抠图（不透明部分），并对边界区域进行模糊处理
func CutImageWithBorderBlur(oriImage image.Image, templateImage image.Image, xPos, yPos, blurSize int) image.Image {
	templateBounds := templateImage.Bounds()
	targetImage := image.NewRGBA(templateBounds)

	// 处理内部像素，与之前相同
	for y := 0; y < templateBounds.Dy(); y++ {
		for x := 0; x < templateBounds.Dx(); x++ {
			alpha := templateImage.At(x, y).(color.NRGBA).A
			if alpha > 100 {
				bgRgb := oriImage.At(xPos+x, yPos+y)
				targetImage.Set(x, y, bgRgb)
			}
		}
	}

	// 处理边界像素，对边界区域进行模糊处理
	for y := 0; y < templateBounds.Dy(); y++ {
		for x := 0; x < templateBounds.Dx(); x++ {
			alpha := templateImage.At(x, y).(color.NRGBA).A
			if alpha > 100 {
				if x < blurSize || x >= templateBounds.Dx()-blurSize || y < blurSize || y >= templateBounds.Dy()-blurSize {
					// 在边界区域，对像素进行模糊处理
					VagueImage(targetImage, x, y)
				}
			}
		}
	}

	return targetImage
}

// VagueImage 对指定坐标的像素进行模糊处理
func VagueImage(img *image.RGBA, x, y int) {
	var red uint32
	var green uint32
	var blue uint32
	var alpha uint32
	var count uint32

	points := [8][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}, {1, 1}, {1, -1}, {-1, 1}, {-1, -1}}

	for _, point := range points {
		pointX := x + point[0]
		pointY := y + point[1]

		if pointX < 0 || pointX >= img.Bounds().Dx() || pointY < 0 || pointY >= img.Bounds().Dy() {
			continue
		}

		r, g, b, a := img.RGBAAt(pointX, pointY).RGBA()
		red += r >> 8
		green += g >> 8
		blue += b >> 8
		alpha += a >> 8
		count++
	}

	if count == 0 {
		return
	}

	avg := count

	rgba := color.RGBA{R: uint8(red / avg), G: uint8(green / avg), B: uint8(blue / avg), A: uint8(alpha / avg)}

	img.SetRGBA(x, y, rgba)
}

// OverlayImage 图片覆盖（覆盖图压缩到width*height大小，覆盖到底图上）
// baseImage：底图
// coverImage：覆盖图
// x：起始x轴
// y：起始y轴
func OverlayImage(baseImage, coverImage image.Image, x, y int) *image.RGBA {

	// 获取覆盖图的边界
	coverBounds := coverImage.Bounds()

	// 创建一个新的 *image.RGBA，将底图复制到新图像上
	rgba := image.NewRGBA(image.Rect(0, 0, baseImage.Bounds().Dx(), baseImage.Bounds().Dy()))
	draw.Draw(rgba, baseImage.Bounds(), baseImage, image.Point{}, draw.Over)

	// 计算覆盖图的位置和大小
	coverRect := image.Rect(x, y, x+coverBounds.Dx(), y+coverBounds.Dy())

	// 将覆盖图绘制到新图像上
	draw.Draw(rgba, coverRect, coverImage, image.Point{}, draw.Over)

	return rgba
}

// CreateTransparentImage 创建一个透明图像
// width：图像宽度
// height：图像高度
func CreateTransparentImage(width, height int) image.Image {
	// 创建一个透明图像
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	return img
}

// ConvertToRGBA 将 image.Image 转换为 *image.RGBA
func ConvertToRGBA(input image.Image) *image.RGBA {
	// 获取输入图像的边界信息
	bounds := input.Bounds()

	// 创建一个新的 *image.RGBA
	output := image.NewRGBA(bounds)

	// 使用 draw 包将像素从输入图像复制到输出图像
	draw.Draw(output, bounds, input, bounds.Min, draw.Src)

	return output
}

// ToBase64 将 image.Image 转换为 Base64 字符串
func ToBase64(img image.Image) (string, error) {
	// 创建一个字节数组缓冲区
	var buf bytes.Buffer

	// 将图像编码为PNG格式并写入缓冲区
	if err := png.Encode(&buf, img); err != nil {
		return "", err
	}

	// 将字节数组转换为Base64字符串
	encoded := base64.StdEncoding.EncodeToString(buf.Bytes())

	return encoded, nil
}

func ModifyImageRGBA(img image.Image, rgba color.RGBA) image.Image {
	bounds := img.Bounds()
	width, height := bounds.Dx(), bounds.Dy()

	var newImg *image.RGBA

	if imgRGBA, ok := img.(*image.RGBA); ok {
		newImg = image.NewRGBA(image.Rect(0, 0, width, height))
		copy(newImg.Pix, imgRGBA.Pix)
	} else if imgNRGBA, ok := img.(*image.NRGBA); ok {
		newImg = image.NewRGBA(image.Rect(0, 0, width, height))
		copy(newImg.Pix, imgNRGBA.Pix)
	} else {
		log.Fatal("Unsupported image format")
	}

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			// 获取像素颜色
			r, g, b, a := newImg.At(x, y).RGBA()

			// 如果是黑色，将其更改为新颜色
			if r == 0 && g == 0 && b == 0 && a > 0 {
				newImg.Set(x, y, rgba)
			}
		}
	}

	return newImg
}
