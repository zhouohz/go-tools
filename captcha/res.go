package captcha

import (
	"github.com/fogleman/gg"
	"github.com/zhouohz/go-tools/core/file"
	image2 "github.com/zhouohz/go-tools/core/image"
	"github.com/zhouohz/go-tools/core/util/random"
	"image"
	"image/color"
	"log"
	"math/rand"
	"time"
	"unicode"
)

var innerBg []image.Image
var innerFont []string

const bgDir = "captcha/resources/bg"

const fontDir = "captcha/resources/font"

func init() {
	suffix1, _ := file.FileListBySuffix(bgDir, ".jpg", true, false, 0)
	innerBg = loadBgByNames(suffix1)

	suffix2, _ := file.FileListBySuffix(fontDir, ".ttf", true, false, 0)
	innerFont = suffix2
}

func loadBgByNames(assetNames []string) []image.Image {
	fonts := make([]image.Image, 0)
	for _, assetName := range assetNames {
		f := loadBgByName(assetName)
		fonts = append(fonts, f)
	}
	return fonts
}

func loadBgByName(name string) image.Image {
	img, err := image2.ParseImage(name)
	if err != nil {
		panic(err)
	}

	return img
}

func RandGetBg() image.Image {
	return innerBg[random.RandInt(len(innerBg))]
}

func RandGetFont() string {
	return innerFont[random.RandInt(len(innerFont))]
}

func RandomColor() color.Color {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return color.RGBA{R: uint8(r.Intn(256)), G: uint8(r.Intn(256)), B: uint8(r.Intn(256)), A: 255}
}

// SetWatermark adds a watermark to an image.
func SetWatermark(img image.Image, text string, fontsize float64, co color.RGBA) (image.Image, error) {
	width := img.Bounds().Dx()
	height := img.Bounds().Dy()

	// 创建一个透明的RGBA图像
	imgb := image.NewRGBA(image.Rect(0, 0, width, height))

	dc := gg.NewContextForRGBA(imgb)
	dc.SetColor(color.Transparent) // 设置背景颜色为透明
	dc.Clear()

	dc.SetColor(co)

	// 加载汉字字体
	err := dc.LoadFontFace(RandGetFont(), fontsize) // 替换成你的汉字字体文件路径
	if err != nil {
		log.Fatal(err)
	}

	// 计算文本的位置，放在右下角
	textWidth, textHeight := dc.MeasureString(text)
	x := float64(width) - textWidth - 10 // Adjust the right margin (e.g., 10 pixels)
	y := float64(height) - textHeight    // Adjust the bottom margin (e.g., 10 pixels)

	// 在指定位置绘制汉字文本
	dc.RotateAbout(0, float64(width)/2.0, float64(height)/2.0)
	dc.DrawStringAnchored(text, x, y, 0, 0)

	return image2.OverlayImage(img, dc.Image(), 0, 0), nil
}

func GetEnOrChLength(text string) int {
	enCount, zhCount := 0, 0

	for _, t := range text {
		if unicode.Is(unicode.Han, t) {
			zhCount++
		} else {
			enCount++
		}
	}

	chOffset := (25/2)*zhCount + 5
	enOffset := enCount * 8

	return chOffset + enOffset
}

// RandomLightColor 生成一个明显可见的随机颜色
func RandomLightColor() color.Color {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// 限制颜色的亮度（Value）和饱和度（Saturation），以确保颜色更明显可见
	minValue := 0.7
	minSaturation := 0.5

	hue := r.Float64() * 360.0
	saturation := minSaturation + r.Float64()*(1.0-minSaturation)
	value := minValue + r.Float64()*(1.0-minValue)

	return hsvToRGB(hue, saturation, value)
}

// hsvToRGB 将 HSV 转换为 RGB
func hsvToRGB(h, s, v float64) color.RGBA {
	h = h / 60.0
	i := int(h)
	f := h - float64(i)
	p := v * (1 - s)
	q := v * (1 - f*s)
	t := v * (1 - (1-f)*s)

	switch i {
	case 0:
		return color.RGBA{uint8(v * 255), uint8(t * 255), uint8(p * 255), 255}
	case 1:
		return color.RGBA{uint8(q * 255), uint8(v * 255), uint8(p * 255), 255}
	case 2:
		return color.RGBA{uint8(p * 255), uint8(v * 255), uint8(t * 255), 255}
	case 3:
		return color.RGBA{uint8(p * 255), uint8(q * 255), uint8(v * 255), 255}
	case 4:
		return color.RGBA{uint8(t * 255), uint8(p * 255), uint8(v * 255), 255}
	default:
		return color.RGBA{uint8(v * 255), uint8(p * 255), uint8(q * 255), 255}
	}
}
