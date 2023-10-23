package slide

import (
	image2 "github.com/zhouohz/go-tools/core/image"
	"github.com/zhouohz/go-tools/core/util/random"
	"golang.org/x/image/colornames"
	"golang.org/x/image/draw"
	"image"
	"image/png"
	"log"
	"os"
	"testing"
)

func TestName(t *testing.T) {
	background := "./res/bg/1.png"
	//active := "active.png"
	fixed := "./res/block/1.png"
	backImage, err := image2.ParseImage(background)
	if err != nil {
		log.Panic(err)
	}
	fixedImage, err := image2.ParseImage(fixed)
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

	cutImage := image2.CutImage(backImage, fixedImage, randomX, randomY)

	Cut(backImage, cutImage, 100, 100)
	//backImage = OverlayImage(backImage, fixedImage, randomX, randomY)
	//cutImage = OverlayImage(cutImage, actImage, 0, 0)
	//
	//matrixTemplate := CreateTransparentImage(actImage.Bounds().Dx(), backImage.Bounds().Dy())
	////
	//i := OverlayImage(matrixTemplate, cutImage, 0, randomY)

	//if err := SaveAs(i, "cut.jpg"); err != nil {
	//	log.Panic(err)
	//}
	if err := SaveAs(cutImage, "bg.jpg"); err != nil {
		log.Panic(err)
	}

	if err := SaveAs(backImage, "bg1.jpg"); err != nil {
		log.Panic(err)
	}
}

func Cut(backgroundImage, templateImage image.Image, x1, y1 int) {
	xLength := templateImage.Bounds().Dx()
	yLength := templateImage.Bounds().Dy()

	for x := 0; x < xLength; x++ {
		for y := 0; y < yLength; y++ {
			// 如果模板图像当前像素点不是透明色 copy源文件信息到目标图片中
			isOpacity := IsOpacity(templateImage, x, y)
			// 当前模板像素在背景图中的位置
			backgroundX := x + x1
			backgroundY := y + y1

			//// 当不为透明时
			//if !isOpacity {
			//	// 获取原图像素
			//	backgroundRgba := ImageToRGBA(backgroundImage).RGBAAt(backgroundX, backgroundY)
			//	// 将原图的像素扣到模板图上
			//	templateImage.SetPixel(backgroundRgba, x, y)
			//	// 背景图区域模糊
			//	backgroundImage.VagueImage(backgroundX, backgroundY)
			//}

			//防止数组越界判断
			if x == (xLength-1) || y == (yLength-1) {
				continue
			}

			rightOpacity := IsOpacity(templateImage, x+1, y)
			downOpacity := IsOpacity(templateImage, x, y+1)

			//描边处理，,取带像素和无像素的界点，判断该点是不是临界轮廓点,如果是设置该坐标像素是白色
			if (isOpacity && !rightOpacity) || (!isOpacity && rightOpacity) || (isOpacity && !downOpacity) || (!isOpacity && downOpacity) {
				ImageToRGBA(backgroundImage).SetRGBA(x, y, colornames.White)
				ImageToRGBA(backgroundImage).SetRGBA(backgroundX, backgroundY, colornames.White)
			}
		}
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
