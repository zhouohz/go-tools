package main

import (
	"fmt"
	"github.com/fogleman/gg"
	"github.com/zhouohz/go-tools/captcha/core/puzzle"
	"github.com/zhouohz/go-tools/captcha/core/seq"
	"github.com/zhouohz/go-tools/captcha/resources/click"
	image2 "github.com/zhouohz/go-tools/core/image"
	"github.com/zhouohz/go-tools/core/util/random"
	"image"
	"image/color"
	"log"
	"math"
	"math/rand"
	"testing"
	"time"
)

//func TestName(t *testing.T) {
//	//随机背景
//	//imgPath := "/captcha/resources/click/2.jpg"
//	//image, err := image2.ParseImage(imgPath)
//	//if err != nil {
//	//	panic(err)
//	//}
//	////随机字
//	num := 5
//	fmt.Println(getRandomLetters(num))
//	//随机font
//	//随机位置
//	//avg := image.Bounds().Dx() / num
//	// 随机x
//	//for i := 0; i < num; i++ {
//	//	randomX := 0
//	//	if i == 0 {
//	//		randomX = 1
//	//	} else {
//	//		randomX = avg * i
//	//	}
//	//
//	//	// 随机y
//	//	//randomY := random.RandomIntRange(10, bgImage.getHeight()-clickImgHeight)
//	//}
//
//	//wordImage := drawTextOnTransparentBackground("周", getRandomColor())
//
//	//image2.SaveAs(wordImage, "p.png")
//
//}

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

const (
	width    = 56
	height   = 56
	fontSize = 48
	text     = "人"
	fontPath = "captcha/resources/click/font/ZCOOLQingKeHuangYou-Regular.ttf" // 替换成你的字体文件路径
)

type Letter struct {
	point  point
	letter string
}

type point struct {
	x, y int
}

func TestImage1(t *testing.T) {
	//随机背景
	imgPath := "captcha/resources/click/bg/2.jpg"
	image, err := image2.ParseImage(imgPath)
	if err != nil {
		panic(err)
	}
	num := 4
	letters := getRandomLetters(num)
	avg := image.Bounds().Dx() / num
	ls := make([]Letter, num)
	for i := range letters {
		img := textImage(string(letters[i]), getRandomColor(), random.RandomFloatRange(0.0, 360.0))
		// 随机x
		randomX := 0
		if i == 0 {
			randomX = 10
		} else {
			randomX = avg * i
		}
		// 随机y
		randomY := random.RandomIntRange(10, image.Bounds().Dy()-img.Bounds().Dy())
		image = image2.OverlayImage(image, img, randomX, randomY)
		ls[i] = Letter{
			point:  point{x: randomX, y: randomY},
			letter: string(letters[i]),
		}
	}

	randomLetters := RandomLetters(ls, 3)

	fmt.Println(randomLetters)

	image2.SaveAs(image, "output.jpg")

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

func TestPuzzle(t *testing.T) {
	get := puzzle.Get()

	fmt.Println(get)
}

func TestName(t *testing.T) {
	se := seq.New()

	if err := se.LoadWordDict("captcha/core/seq/dict.txt"); err != nil {
		panic(err)
	}

	_, letters := se.Get()

	fmt.Println(letters)
}

func hasDuplicateChars(s string) bool {
	charCount := make(map[rune]int)

	for _, char := range s {
		charCount[char]++

		// 如果字符出现次数超过一次，返回 true
		if charCount[char] > 1 {
			return true
		}
	}

	// 如果遍历完成没有找到相同字符，返回 false
	return false
}
