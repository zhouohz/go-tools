package click

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/fogleman/gg"
	"github.com/zhouohz/go-tools/captcha"
	"github.com/zhouohz/go-tools/captcha/resources/click"
	"github.com/zhouohz/go-tools/captcha/store"
	image2 "github.com/zhouohz/go-tools/core/image"
	"github.com/zhouohz/go-tools/core/util/id"
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
)

type ClickLetter struct {
	num    int //文字数量
	offset int //偏移量

	store store.Cache
}

type Letter struct {
	Point  Point
	Letter string
}

type Point struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type Letters []Letter

func (ls Letters) Letter() []string {
	letters := make([]string, len(ls))
	for i, l := range ls {
		letters[i] = l.Letter
	}
	return letters
}

func (ls Letters) Points() []Point {
	points := make([]Point, len(ls))
	for i, l := range ls {
		points[i] = l.Point
	}
	return points
}

func New(num, offset int, cache store.Cache) *ClickLetter {
	return &ClickLetter{
		num:    num,
		offset: offset,
		store:  cache,
	}
}

func (this *ClickLetter) Get(ctx context.Context) (*captcha.Generate, error) {
	//随机背景
	bgImg := captcha.RandGetBg()

	bgImg, err := captcha.SetWatermark(bgImg, "寰宇天穹", 18, color.RGBA{R: 255, G: 255, B: 255, A: 195})
	if err != nil {
		return nil, err
	}

	letters := this.getRandomLetters(this.num)
	bgImg, generate, err := this.generate(bgImg, letters)
	if err != nil {
		return nil, err
	}

	points := generate.Points()

	uuid := id.IdUtil().FastSimpleUUID()

	base64, err := image2.ToBase64(bgImg, true)
	if err != nil {
		return nil, err
	}
	// 将 points 转换为 JSON 字符串
	pointsJSON, err := json.Marshal(points)
	if err != nil {
		return nil, err
	}
	this.store.Set(ctx, captcha.GetID(uuid), string(pointsJSON), 300)

	return &captcha.Generate{
		Bg:     base64,
		Front:  generate.Letter(),
		Token:  uuid,
		Answer: points,
	}, nil
}

func (this *ClickLetter) Check(ctx context.Context, token, pointJson string) error {
	var ps []Point
	if err := json.Unmarshal([]byte(pointJson), &ps); err != nil {
		return err
	}

	val := this.store.Get(ctx, captcha.GetID(token))
	var points []Point
	if err := json.Unmarshal([]byte(val), &points); err != nil {
		fmt.Println(err)
		return err
	}
	if len(ps) != len(points) {
		return captcha.CheckCodeError
	}

	for i := range ps {
		if !number.NumInRange(ps[i].X, points[i].X-(fontSize/2), points[i].X+(fontSize/2), true) ||
			!number.NumInRange(ps[i].Y, points[i].Y-(fontSize/2), points[i].Y+(fontSize/2), true) {
			return captcha.CheckCodeError
		}

	}

	return nil
}

func (this *ClickLetter) generate(bgImg image.Image, letters []rune) (image.Image, Letters, error) {
	interval := number.CalculateHypotenuse(width, height)
	avg := bgImg.Bounds().Dx() / interval
	num := len(letters)
	repair := (bgImg.Bounds().Dx() - (interval * num)) / (num + 1)
	if avg < num {
		return nil, nil, errors.New("too many letters")
	}
	ls := make([]Letter, num)
	font := captcha.RandGetFont()
	for i := range letters {
		img := this.textImage(string(letters[i]), font, captcha.RandomColor(), random.RandomFloatRange(0.0, 360.0))
		// 随机x
		randomX := (interval * i) + (interval / 2) + (repair * (i + 1))
		// 随机y
		randomY := random.RandomIntRange(height/2, bgImg.Bounds().Dy()-(img.Bounds().Dy()/2)-30)
		bgImg = image2.OverlayImageAtCenter(bgImg, img, randomX, randomY)
		ls[i] = Letter{
			Point:  Point{X: randomX, Y: randomY},
			Letter: string(letters[i]),
		}
	}

	return bgImg, this.randomLetters(ls, num-1), nil
}

func (this *ClickLetter) randomLetters(objects []Letter, count int) Letters {
	rand.Shuffle(len(objects), func(i, j int) {
		objects[i], objects[j] = objects[j], objects[i]
	})
	return objects[:count]
}

func (this *ClickLetter) toRadians(angdeg float64) float64 {
	return angdeg * (math.Pi / 180.0)
}

func (this *ClickLetter) textImage(text string, f string, co color.Color, d float64) image.Image {
	// 创建一个透明的RGBA图像
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	dc := gg.NewContextForRGBA(img)
	dc.SetColor(color.Transparent) // 设置背景颜色为透明
	dc.Clear()

	dc.SetColor(co)

	// 加载汉字字体
	err := dc.LoadFontFace(f, fontSize) // 替换成你的汉字字体文件路径
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

func (this *ClickLetter) getRandomLetters(num int) []rune {
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
