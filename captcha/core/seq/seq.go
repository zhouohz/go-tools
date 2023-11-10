package seq

import (
	"bufio"
	"fmt"
	"github.com/fogleman/gg"
	"github.com/zhouohz/go-tools/captcha"
	"github.com/zhouohz/go-tools/core/collection/list"
	image2 "github.com/zhouohz/go-tools/core/image"
	"github.com/zhouohz/go-tools/core/util/id"
	"github.com/zhouohz/go-tools/core/util/number"
	"github.com/zhouohz/go-tools/core/util/random"
	"image"
	"image/color"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"
)

const (
	width    = 40
	height   = 40
	fontSize = 36
	text     = "人"
	fontPath = "captcha/resources/click/font/ZCOOLQingKeHuangYou-Regular.ttf" // 替换成你的字体文件路径
)

type Seq struct {
	idioms [][]rune
}

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

// New
func New() *Seq {
	return &Seq{
		idioms: make([][]rune, 0),
	}
}

func (this *Seq) Get() (captcha.GetRes, error) {
	//随机背景
	imgPath := fmt.Sprintf("captcha/resources/seq/bg/%d.jpg", random.RandomIntRange(1, 7))
	image, err := image2.ParseImage(imgPath)
	if err != nil {
		panic(err)
	}
	num := 4
	letters := this.getRandomIdiom()
	interval := number.CalculateHypotenuse(width, height)
	avg := image.Bounds().Dx() / interval
	repair := (image.Bounds().Dx() - (interval * num)) / (num + 1)
	if avg < num {
		panic("超过数量")
	}

	ls := make(Letters, num)
	shuffle := list.Shuffle([]int{0, 1, 2, 3})
	for i := range letters {
		img := this.textImage(string(letters[i]), this.getRandomColor(), random.RandomFloatRange(0.0, 360.0))
		// 随机x
		randomX := (interval * shuffle[i]) + (interval / 2) + (repair * (shuffle[i] + 1))
		// 随机y
		randomY := random.RandomIntRange(height/2, image.Bounds().Dy()-(img.Bounds().Dy()/2))
		image = image2.OverlayImageAtCenter(image, img, randomX, randomY)
		ls[i] = Letter{
			Point:  Point{X: randomX, Y: randomY},
			Letter: string(letters[i]),
		}
	}

	return captcha.GetRes{
		Bg:    image,
		Front: 4,
		Token: id.IdUtil().FastSimpleUUID(),
		Type:  4,
	}, nil

}

func (this *Seq) Check(checkReq captcha.CheckReq) (captcha.CheckRsp, error) {

}

func (this *Seq) randomLetters(objects []rune) []rune {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(objects), func(i, j int) {
		objects[i], objects[j] = objects[j], objects[i]
	})
	return objects
}

func (this *Seq) textImage(text string, co color.Color, d float64) image.Image {
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
	dc.RotateAbout(this.toRadians(d), float64(img.Bounds().Dx())/2.0, float64(img.Bounds().Dy())/2.0)
	dc.DrawStringAnchored(text, x, y, 0, 0)

	return dc.Image()
}

func (this *Seq) toRadians(angdeg float64) float64 {
	return angdeg * (math.Pi / 180.0)
}

// 定义随机颜色
func (this *Seq) getRandomColor() color.Color {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return color.RGBA{R: uint8(r.Intn(256)), G: uint8(r.Intn(256)), B: uint8(r.Intn(256)), A: 255}
}

func (this *Seq) getRandomIdiom() []rune {

	if len(this.idioms) == 0 {
		panic("please provide idioms")
	}

	// 创建一个新的随机数生成器，以避免并发问题
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	return this.idioms[r.Intn(len(this.idioms))]
}

// LoadWordDict
func (this *Seq) LoadWordDict(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	return this.Load(f)
}

// LoadNetWordDict
func (this *Seq) LoadNetWordDict(url string) error {
	c := http.Client{
		Timeout: 5 * time.Second,
	}
	rsp, err := c.Get(url)
	if err != nil {
		return err
	}
	defer rsp.Body.Close()

	return this.Load(rsp.Body)
}

// Load common method to add words
func (this *Seq) Load(rd io.Reader) error {
	buf := bufio.NewReader(rd)
	for {
		line, _, err := buf.ReadLine()
		if err != nil {
			if err != io.EOF {
				return err
			}
			break
		}

		// 去除空白字符，检查是否为空字符串，如果不是空字符串再添加到 this.idioms 切片中
		cleanedLine := strings.TrimSpace(string(line))
		if cleanedLine != "" {
			this.idioms = append(this.idioms, []rune(cleanedLine))
		}
	}

	return nil
}
