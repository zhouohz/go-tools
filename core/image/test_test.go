package image

import (
	"github.com/zhouohz/go-tools/core/util/random"
	"image"
	"image/png"
	"log"
	"os"
	"testing"
)

func TestOverlayImage(t *testing.T) {
	background := "1.jpg"
	//active := "active.png"
	fixed := "1.png"
	backImage, err := ParseImage(background)
	if err != nil {
		log.Panic(err)
	}
	fixedImage, err := ParseImage(fixed)
	if err != nil {
		log.Panic(err)
	}
	//actImage, err := ParseImage(active)
	//if err != nil {
	//	log.Panic(err)
	//}

	// 获取随机的 x 和 y 轴
	randomX := random.RandomIntRange(fixedImage.Bounds().Dx()+5, backImage.Bounds().Dx()-fixedImage.Bounds().Dx()-10)
	randomY := random.RandInt(backImage.Bounds().Dy() - fixedImage.Bounds().Dy())

	cutImage := CutImage(backImage, fixedImage, randomX, randomY)
	//backImage = OverlayImage(backImage, fixedImage, randomX, randomY)
	//cutImage = OverlayImage(cutImage, actImage, 0, 0)
	//
	//matrixTemplate := CreateTransparentImage(actImage.Bounds().Dx(), backImage.Bounds().Dy())
	////
	//i := OverlayImage(matrixTemplate, cutImage, 0, randomY)

	//if err := SaveAs(i, "cut.jpg"); err != nil {
	//	log.Panic(err)
	//}
	if err := SaveAs(backImage, "bg.jpg"); err != nil {
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
