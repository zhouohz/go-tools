package seq

import (
	"bufio"
	"context"
	"encoding/json"
	"github.com/fogleman/gg"
	"github.com/zhouohz/go-tools/captcha"
	"github.com/zhouohz/go-tools/captcha/store"
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
	width       = 40
	height      = 40
	fontSize    = 36
	defaultDict = "captcha/resources/seq/dict.txt"
)

var temp = make(map[string][]Point)

type Seq struct {
	idioms [][]rune
	offset int
	store  store.Cache
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

func (ls Letters) Points() []Point {
	points := make([]Point, len(ls))
	for i, l := range ls {
		points[i] = l.Point
	}
	return points
}

// New
func New(offset int, store store.Cache) *Seq {
	seq := Seq{
		idioms: make([][]rune, 0),
		store:  store,
		offset: offset,
	}
	seq.LoadWordDict(defaultDict)

	return &seq
}

func (this *Seq) Get(ctx context.Context) (*captcha.Generate, error) {
	//随机背景
	bgImg := captcha.RandGetBg()

	bgImg, err := captcha.SetWatermark(bgImg, "寰宇天穹", 18, color.RGBA{R: 255, G: 255, B: 255, A: 195})
	if err != nil {
		return nil, err
	}

	letters := this.getRandomIdiom()

	bgImg, ls := this.generate(bgImg, letters)
	points := ls.Points()

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
		Bg:    base64,
		Front: 4,
		Token: uuid,
	}, nil

}

func (this *Seq) Check(ctx context.Context, token, pointJson string) error {
	var ps []Point
	if err := json.Unmarshal([]byte(pointJson), &ps); err != nil {
		return err
	}
	val := this.store.Get(ctx, captcha.GetID(token))
	var points []Point
	if err := json.Unmarshal([]byte(val), &points); err != nil {
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

func (this *Seq) generate(bgImg image.Image, letters []rune) (image.Image, Letters) {
	interval := number.CalculateHypotenuse(width, height)
	repair := (bgImg.Bounds().Dx() - (interval * 4)) / 5
	ls := make(Letters, 4)
	shuffle := list.Shuffle([]int{0, 1, 2, 3})
	font := captcha.RandGetFont()
	for i := range letters {
		img := this.textImage(string(letters[i]), font, captcha.RandomColor(), random.RandomFloatRange(0.0, 360.0))
		// 随机x
		randomX := (interval * shuffle[i]) + (interval / 2) + (repair * (shuffle[i] + 1))
		// 随机y
		randomY := random.RandomIntRange(height/2, bgImg.Bounds().Dy()-(img.Bounds().Dy()/2)-30)
		bgImg = image2.OverlayImageAtCenter(bgImg, img, randomX, randomY)
		ls[i] = Letter{
			Point:  Point{X: randomX, Y: randomY},
			Letter: string(letters[i]),
		}
	}

	return bgImg, ls
}

func (this *Seq) randomLetters(objects []rune) []rune {
	rand.Shuffle(len(objects), func(i, j int) {
		objects[i], objects[j] = objects[j], objects[i]
	})
	return objects
}

func (this *Seq) textImage(text string, f string, co color.Color, d float64) image.Image {
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

func (this *Seq) toRadians(angdeg float64) float64 {
	return angdeg * (math.Pi / 180.0)
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
