package slide

import (
	"context"
	"encoding/json"
	"github.com/zhouohz/go-tools/captcha"
	"github.com/zhouohz/go-tools/captcha/store"
	image2 "github.com/zhouohz/go-tools/core/image"
	"github.com/zhouohz/go-tools/core/util/id"
	"github.com/zhouohz/go-tools/core/util/number"
	"github.com/zhouohz/go-tools/core/util/random"
	"image"
	"image/color"
)

// Slide captcha config for captcha-engine-slide.
type Slide struct {
	offset int //容错偏移量

	bg, active, fixed image.Image

	store store.Cache
}

type point struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type Res struct {
	Key        string
	BgImage    *img
	BlockImage *img
}

type img struct {
	Width  int    `json:"width"`
	Height int    `json:"height"`
	Base64 string `json:"base64"`
}

// New creates driver
func New(offset int, store store.Cache) *Slide {
	return &Slide{offset: offset, store: store}
}

func (this *Slide) Get(ctx context.Context) (*captcha.Generate, error) {
	// 获取随机block和bg
	block := RandGetBlock()
	bg := captcha.RandGetBg()

	bg, err := captcha.SetWatermark(bg, "寰宇天穹", 18, color.RGBA{R: 255, G: 255, B: 255, A: 195})
	if err != nil {
		return nil, err
	}

	//获取
	point, bgImage, blockImage := this.generateImage(bg, block)

	//存储图片信息，用于轨迹校验

	uuid := id.IdUtil().FastSimpleUUID()

	base641, err := image2.ToBase64(bgImage, true)
	if err != nil {
		return nil, err
	}
	base642, err := image2.ToBase64(blockImage, true)
	if err != nil {
		return nil, err
	}

	// 将 points 转换为 JSON 字符串
	pointsJSON, err := json.Marshal(point)
	if err != nil {
		return nil, err
	}
	this.store.Set(ctx, captcha.GetID(uuid), string(pointsJSON), 300)
	return &captcha.Generate{
		Bg:    base641,
		Front: base642,
		Token: uuid,
		//Answer: point,
	}, nil

}

func (this *Slide) Check(ctx context.Context, token, pointJson string) error {
	var ps point
	if err := json.Unmarshal([]byte(pointJson), &ps); err != nil {
		return err
	}
	val := this.store.Get(ctx, captcha.GetID(token))
	var p point
	if err := json.Unmarshal([]byte(val), &p); err != nil {
		return err
	}

	if !number.NumInRange(ps.X, p.X-this.offset, p.X+this.offset, true) {
		return captcha.CheckCodeError
	}

	return nil
}

// generatePoint 生成坐标点
func (this *Slide) generatePoint(bg, fixed image.Image) *point {
	randomX := random.RandomIntRange(fixed.Bounds().Dx()+5, bg.Bounds().Dx()-fixed.Bounds().Dx()-10)
	randomY := random.RandInt(bg.Bounds().Dy() - fixed.Bounds().Dy() - 30)
	return &point{X: randomX, Y: randomY}
}

func (this *Slide) generateImage(bg image.Image, block Block) (*point, *image.RGBA, *image.RGBA) {

	generatePoint := this.generatePoint(bg, block.Fixed)
	backOverlayImage := image2.OverlayImage(bg, block.Fixed, generatePoint.X, generatePoint.Y)
	cutImage := image2.CutImage(bg, block.Fixed, generatePoint.X, generatePoint.Y)
	overlayImage := image2.OverlayImage(block.Active, cutImage, 0, 0)
	matrixTemplate := image2.CreateTransparentImage(block.Active.Bounds().Dx(), bg.Bounds().Dy())
	blockImage := image2.OverlayImage(matrixTemplate, overlayImage, 0, generatePoint.Y)

	return generatePoint, backOverlayImage, blockImage
}

func (this *Slide) generateInterference(bg image.Image, block Block, p *point) *image.RGBA {
	// 获取背景图像的宽度和高度
	bgWidth := bg.Bounds().Dx()
	bgHeight := bg.Bounds().Dy()

	// 获取缺口图像的宽度和高度
	blockWidth := block.Fixed.Bounds().Dx()
	blockHeight := block.Fixed.Bounds().Dy()

	yRandomInternal := make([]random.Interval, 0)
	// 如果fix上面有一个fix的区域，就要上面
	if p.Y+(2*blockHeight) < bgHeight {
		yRandomInternal = append(yRandomInternal, random.Interval{
			Start: p.Y + blockHeight,
			End:   bgHeight - blockHeight,
		})
	}
	// 如果fix下面有一个fix的区域，就要上面
	if p.Y > blockHeight {
		yRandomInternal = append(yRandomInternal, random.Interval{
			Start: 0,
			End:   p.Y - blockHeight,
		})
	}
	if len(yRandomInternal) != 0 {
		randomX := random.RandomIntRange(blockWidth+5, bgWidth-blockWidth-10)
		randomY := random.RandNumberInIntervals(yRandomInternal)
		return image2.OverlayImage(bg, block.Fixed, randomX, randomY)
	}
	// 返回带有干扰位置的图像
	return nil
}
