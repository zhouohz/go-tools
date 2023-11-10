package puzzle

import (
	"fmt"
	image2 "github.com/zhouohz/go-tools/core/image"
	"github.com/zhouohz/go-tools/core/util/random"
	"golang.org/x/image/draw"
	"image"
	"math/rand"
	"time"
)

func Get() []image.Image {
	//随机背景
	imgPath := fmt.Sprintf("captcha/resources/click/bg/%d.jpg", random.RandomIntRange(1, 7))
	bg, err := image2.ParseImage(imgPath)
	if err != nil {
		panic(err)
	}
	return ShuffleAndRandomSwap(SplitImage(bg, 2, 4), 2)
	//

	//num := 4
	//letters := getRandomLetters(num)
	//interval := number.CalculateHypotenuse(width, height)
	//avg := image.Bounds().Dx() / interval
	//repair := (image.Bounds().Dx() - (interval * num)) / (num + 1)
	//if avg < num {
	//	panic("超过数量")
	//}
	//ls := make([]Letter, num)
	//for i := range letters {
	//	img := textImage(string(letters[i]), getRandomColor(), random.RandomFloatRange(0.0, 360.0))
	//	// 随机x
	//	randomX := (interval * i) + (interval / 2) + (repair * (i + 1))
	//	// 随机y
	//	randomY := random.RandomIntRange(height/2, image.Bounds().Dy()-(img.Bounds().Dy()/2))
	//	image = image2.OverlayImageAtCenter(image, img, randomX, randomY)
	//	ls[i] = Letter{
	//		Point:  Point{X: randomX, Y: randomY},
	//		Letter: string(letters[i]),
	//	}
	//}
	//
	//return image, RandomLetters(ls, num-1)
}

func ShuffleAndRandomSwap[T comparable](array []T, num int) []T {
	if len(array) < num {
		num = len(array)
	}

	source := rand.NewSource(time.Now().UnixNano()) // 使用当前时间作为随机种子
	rng := rand.New(source)

	// 随机选择两个不同的索引
	index1 := rng.Intn(len(array))
	index2 := rng.Intn(len(array))
	for index2 == index1 {
		index2 = rng.Intn(len(array))
	}

	// 交换这两个随机索引处的元素
	array[index1], array[index2] = array[index2], array[index1]

	return array
}

// SplitImage 分割一张图片为 8 个小块并保存
func SplitImage(img image.Image, hNum, wNum int) []image.Image {

	// 获取图片的宽度和高度
	width := img.Bounds().Dx()
	height := img.Bounds().Dy()

	// 计算每个小块的宽度和高度
	blockWidth := width / wNum
	blockHeight := height / hNum

	images := make([]image.Image, 0)
	for row := 0; row < hNum; row++ {
		for col := 0; col < wNum; col++ {
			x := col * blockWidth
			y := row * blockHeight
			blockRect := image.Rect(x, y, x+blockWidth, y+blockHeight)
			block := image.NewRGBA(blockRect)
			draw.Draw(block, block.Bounds(), img, blockRect.Min, draw.Src)

			images = append(images, block)
		}
	}

	return images
}
