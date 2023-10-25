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

func TestName(t *testing.T) {

	bint := random.RandomIntRange(1, 5)
	blockInt := random.RandomIntRange(1, 3)

	background := fmt.Sprintf("./res/bg/%d.jpg", bint)
	active := fmt.Sprintf("./res/block/%d/active.png", blockInt)
	fixed := fmt.Sprintf("./res/block/%d/fixed.png", blockInt)
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

//func TestGet(t *testing.T) {
//	// 打开图像文件
//	imageFile, err := os.Open("./res/block/1/active.png")
//	if err != nil {
//		fmt.Println("Error opening image:", err)
//		return
//	}
//	defer imageFile.Close()
//
//	// 解码图像
//	img, _, err := image.Decode(imageFile)
//	if err != nil {
//		fmt.Println("Error decoding image:", err)
//		return
//	}
//
//	// 创建一个新的RGBA图像
//	bounds := img.Bounds()
//	newImage := image.NewRGBA(bounds)
//
//	// 遍历图像像素
//	for x := bounds.Min.X; x < bounds.Max.X; x++ {
//		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
//			// 获取当前像素的颜色
//			curColor := img.At(x, y)
//
//			// 检查是否是不透明的像素
//			_, _, _, alpha := curColor.RGBA()
//
//			if alpha > 0 {
//				// 检查上下左右四个方向的像素
//				left := img.At(x-1, y)
//				right := img.At(x+1, y)
//				up := img.At(x, y-1)
//				down := img.At(x, y+1)
//
//				// 如果至少一个相邻像素是透明的，将当前像素的颜色设置为白色
//				if isTransparent(left) || isTransparent(right) || isTransparent(up) || isTransparent(down) {
//					newImage.Set(x, y, color.White)
//				} else {
//					newImage.Set(x, y, curColor)
//				}
//			}
//		}
//	}
//
//	// 保存处理后的图像
//	outputFile, err := os.Create("output.png")
//	if err != nil {
//		fmt.Println("Error creating output file:", err)
//		return
//	}
//	defer outputFile.Close()
//
//	err = png.Encode(outputFile, newImage)
//	if err != nil {
//		fmt.Println("Error encoding image:", err)
//		return
//	}
//
//	fmt.Println("Image processing complete.")
//}
//
//// 检查颜色是否是完全透明
//func isTransparent(c color.Color) bool {
//	_, _, _, alpha := c.RGBA()
//	return alpha == 0
//}
