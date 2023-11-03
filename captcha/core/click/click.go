package click

import (
	"fmt"
	"github.com/fogleman/gg"
	"github.com/zhouohz/go-tools/captcha/resources/click"
	image2 "github.com/zhouohz/go-tools/core/image"
	"github.com/zhouohz/go-tools/core/util/number"
	"github.com/zhouohz/go-tools/core/util/random"
	"image"
	"image/color"
	"log"
	"math"
	"math/rand"
	"time"
)

const (
	width    = 40
	height   = 40
	fontSize = 36
	text     = "人"
	fontPath = "captcha/resources/click/font/ZCOOLQingKeHuangYou-Regular.ttf" // 替换成你的字体文件路径
)

type Letter struct {
	Point  Point
	Letter string
}

type Point struct {
	X, Y int
}

type Letters []Letter

func (ls Letters) Letter() []string {
	letters := make([]string, len(ls))
	for i, l := range ls {
		letters[i] = l.Letter
	}

	return letters
}

func Get() (image.Image, Letters) {
	//随机背景
	imgPath := fmt.Sprintf("captcha/resources/click/bg/%d.jpg", random.RandomIntRange(1, 7))
	image, err := image2.ParseImage(imgPath)
	if err != nil {
		panic(err)
	}
	num := 4
	letters := getRandomLetters(num)
	interval := number.CalculateHypotenuse(width, height)
	avg := image.Bounds().Dx() / interval
	repair := (image.Bounds().Dx() - (interval * num)) / (num + 1)
	if avg < num {
		panic("超过数量")
	}
	ls := make([]Letter, num)
	for i := range letters {
		img := textImage(string(letters[i]), getRandomColor(), random.RandomFloatRange(0.0, 360.0))
		// 随机x
		randomX := (interval * i) + (interval / 2) + (repair * (i + 1))
		// 随机y
		randomY := random.RandomIntRange(height/2, image.Bounds().Dy()-(img.Bounds().Dy()/2))
		image = image2.OverlayImageAtCenter(image, img, randomX, randomY)
		ls[i] = Letter{
			Point:  Point{X: randomX, Y: randomY},
			Letter: string(letters[i]),
		}
	}

	return image, RandomLetters(ls, num-1)
}

func RandomLetters(objects []Letter, count int) []Letter {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(objects), func(i, j int) {
		objects[i], objects[j] = objects[j], objects[i]
	})
	return objects[:count]
}

func toRadians(angdeg float64) float64 {
	return angdeg * (math.Pi / 180.0)
}

func textImage(text string, co color.Color, d float64) image.Image {
	// 创建一个透明的RGBA图像
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	dc := gg.NewContextForRGBA(img)
	dc.SetColor(color.Transparent) // 设置背景颜色为透明
	dc.Clear()

	dc.SetColor(co)

	// 加载汉字字体
	err := dc.LoadFontFace(fontPath, fontSize) // 替换成你的汉字字体文件路径
	if err != nil {
		log.Fatal(err)
	}

	// 计算文本的位置，居中显示
	textWidth, textHeight := dc.MeasureString(text)
	x := (width - textWidth) / 2
	y := (height + textHeight) / 2
	// 在指定位置绘制汉字文本
	dc.RotateAbout(toRadians(d), float64(img.Bounds().Dx())/2.0, float64(img.Bounds().Dy())/2.0)
	dc.DrawStringAnchored(text, x, y, 0, 0)

	return dc.Image()
}

func getRandomLetters(num int) []rune {
	// 创建一个新的随机数生成器，以避免并发问题
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	RunesArray := click.RunesArray

	// 创建一个切片，表示可用的索引位置
	indices := r.Perm(len(RunesArray))

	// 限制索引不超过数组的长度和指定的数量
	if num > len(RunesArray) {
		num = len(RunesArray)
	}

	// 选择不超过指定数量的不同的字母
	selected := make([]rune, num)
	for i := 0; i < num; i++ {
		selected[i] = RunesArray[indices[i]]
	}

	return selected
}

// 定义随机颜色
func getRandomColor() color.Color {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return color.RGBA{R: uint8(r.Intn(256)), G: uint8(r.Intn(256)), B: uint8(r.Intn(256)), A: 255}
}
