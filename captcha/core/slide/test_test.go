package slide

import (
	"fmt"
	image2 "github.com/zhouohz/go-tools/core/image"
	"github.com/zhouohz/go-tools/core/util/random"
	"golang.org/x/image/draw"
	"image"
	"image/png"
	"log"
	"os"
	"testing"
)

func TestImg(t *testing.T) {
	//slide := NewSlide(10)

	//slide.Get()
}

func TestName(t *testing.T) {

	//bint := random.RandomIntRange(1, 3)
	//blockInt := random.RandomIntRange(1, 3)

	background := fmt.Sprintf("./res/bg/%d.jpg", 5)
	active := fmt.Sprintf("./res/block/%d/active.png", 1)
	fixed := fmt.Sprintf("./res/block/%d/fixed.png", 1)
	backImage, err := image2.ParseImage(background)
	if err != nil {
		log.Panic(err)
	}
	fixedImage, err := image2.ParseImage(fixed)
	if err != nil {
		log.Panic(err)
	}
	actImage, err := image2.ParseImage(active)
	if err != nil {
		log.Panic(err)
	}

	// 获取随机的 x 和 y 轴
	randomX := random.RandomIntRange(fixedImage.Bounds().Dx()+5, backImage.Bounds().Dx()-fixedImage.Bounds().Dx()-10)
	randomY := random.RandInt(backImage.Bounds().Dy() - fixedImage.Bounds().Dy())
	backOverlayImage := image2.OverlayImage(backImage, fixedImage, randomX, randomY)
	cutImage := image2.CutImage(backImage, fixedImage, randomX, randomY)
	overlayImage := image2.OverlayImage(actImage, cutImage, 0, 0)

	matrixTemplate := image2.CreateTransparentImage(actImage.Bounds().Dx(), backImage.Bounds().Dy())
	//
	i := image2.OverlayImage(matrixTemplate, overlayImage, 0, randomY)

	if err := SaveAs(i, "bg.png"); err != nil {
		log.Panic(err)
	}

	if err := SaveAs(backOverlayImage, "bg1.png"); err != nil {
		log.Panic(err)
	}
}

func SaveAs(img image.Image, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	err = png.Encode(file, img)
	return err
}

// IsOpacity 该像素是否透明
func IsOpacity(i image.Image, x, y int) bool {
	A := ImageToRGBA(i).RGBAAt(x, y).A
	if float32(A) <= 125 {
		return true
	}
	return false
}

// ImageToRGBA 图片转rgba
func ImageToRGBA(img image.Image) *image.RGBA {
	// No conversion needed if image is an *image.RGBA.
	if dst, ok := img.(*image.RGBA); ok {
		return dst
	}

	// Use the image/draw package to convert to *image.RGBA.
	b := img.Bounds()
	dst := image.NewRGBA(image.Rect(0, 0, b.Dx(), b.Dy()))
	draw.Draw(dst, dst.Bounds(), img, b.Min, draw.Src)
	return dst
}
