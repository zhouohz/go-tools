package slide

import (
	image2 "github.com/zhouohz/go-tools/core/image"
	"github.com/zhouohz/go-tools/core/util/random"
	"image"
)

// Slide captcha config for captcha-engine-slide.
type Slide struct {
	Offset int //容错偏移量

	bg, active, fixed image.Image
}

type point struct {
	x, y int
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

// NewSlide creates driver
func NewSlide(offset int) *Slide {
	return &Slide{Offset: offset}
}

func (this *Slide) Generate() (*Res, error) {
	// 获取随机block和bg
	block := this.randomBlock()
	bg := this.randomBg()

	//获取
	_, bgImage, blockImage := this.generateImage(bg, block)

	bgImageBase64, err := image2.ToBase64(bgImage)
	if err != nil {
		return nil, err
	}

	blockImageBase64, err := image2.ToBase64(blockImage)
	if err != nil {
		return nil, err
	}
	return &Res{
		Key: random.RandomStr(16),
		BgImage: &img{
			Width:  bgImage.Bounds().Dx(),
			Height: bgImage.Bounds().Dy(),
			Base64: bgImageBase64,
		},
		BlockImage: &img{
			Width:  blockImage.Bounds().Dx(),
			Height: blockImage.Bounds().Dy(),
			Base64: blockImageBase64,
		},
	}, nil

}

func (this *Slide) randomBg() image.Image {
	return InnerBg[random.RandInt(len(InnerBg))]
}

func (this *Slide) randomBlock() Block {

	return InnerBLock[random.RandInt(len(InnerBLock))]
}

// generatePoint 生成坐标点
func (this *Slide) generatePoint(bg, fixed image.Image) *point {
	randomX := random.RandomIntRange(fixed.Bounds().Dx()+5, bg.Bounds().Dx()-fixed.Bounds().Dx()-10)
	randomY := random.RandInt(bg.Bounds().Dy() - fixed.Bounds().Dy())
	return &point{x: randomX, y: randomY}
}

func (this *Slide) generateImage(bg image.Image, block Block) (*point, *image.RGBA, *image.RGBA) {

	generatePoint := this.generatePoint(bg, block.Fixed)

	backOverlayImage := image2.OverlayImage(bg, block.Fixed, generatePoint.x, generatePoint.y)
	cutImage := image2.CutImage(bg, block.Fixed, generatePoint.x, generatePoint.y)
	overlayImage := image2.OverlayImage(block.Active, cutImage, 0, 0)

	matrixTemplate := image2.CreateTransparentImage(block.Active.Bounds().Dx(), bg.Bounds().Dy())
	blockImage := image2.OverlayImage(matrixTemplate, overlayImage, 0, generatePoint.y)

	return generatePoint, backOverlayImage, blockImage
}

//func Get() {
//	//image := GetBackgroundImage()
//	////image.Src
//	//block := GetBlockImage()
//	//fmt.Println(generatePoint(image, block))
//
//}
