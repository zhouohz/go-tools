package puzzle

import (
	"context"
	"encoding/json"
	"github.com/zhouohz/go-tools/captcha"
	"github.com/zhouohz/go-tools/captcha/store"
	image2 "github.com/zhouohz/go-tools/core/image"
	"github.com/zhouohz/go-tools/core/util/id"
	"golang.org/x/image/draw"
	"image"
	"image/color"
	"math/rand"
)

type Puzzle struct {
	store store.Cache
}

// New
func New(store store.Cache) *Puzzle {
	return &Puzzle{store: store}
}

func (this *Puzzle) Get(ctx context.Context) (*captcha.Generate, error) {
	//随机背景
	bg := captcha.RandGetBg()

	bg, err := captcha.SetWatermark(bg, "寰宇天穹", 18, color.RGBA{R: 255, G: 255, B: 255, A: 195})
	if err != nil {
		return nil, err
	}

	processImage, index := this.processImage(bg, 2, 4)
	uuid := id.IdUtil().FastSimpleUUID()

	base64, err := image2.ToBase64(processImage, true)
	if err != nil {
		return nil, err
	}

	// 将 points 转换为 JSON 字符串
	pointsJSON, err := json.Marshal(index)
	if err != nil {
		return nil, err
	}
	this.store.Set(ctx, captcha.GetID(uuid), string(pointsJSON), 300)

	return &captcha.Generate{
		Bg:     base64,
		Front:  2,
		Token:  uuid,
		Answer: index,
	}, nil
}

func (this *Puzzle) Check(ctx context.Context, token, pointJson string) error {

	var ps []int
	if err := json.Unmarshal([]byte(pointJson), &ps); err != nil {
		return err
	}

	//points, ok := temp[token]
	//if !ok {
	//	return captcha.InvalidCodeError
	//}

	//if equal := array.EveryEqual(ps, points); !equal {
	//	return captcha.CheckCodeError
	//}

	return nil

}

func (this *Puzzle) processImage(input image.Image, hNum, wNum int) (image.Image, []int) {

	// 获取图片的宽度和高度
	bounds := input.Bounds()
	width := bounds.Max.X
	height := bounds.Max.Y
	num := hNum * wNum
	// 将图片分成八份
	parts := make([]image.Image, num)
	index := make([]int, num)
	partWidth := width / wNum
	partHeight := height / hNum
	in := 0
	for i := 0; i < hNum; i++ {
		for j := 0; j < wNum; j++ {
			x := j * partWidth
			y := i * partHeight
			index[in] = in
			part := image.NewRGBA(image.Rect(0, 0, partWidth, partHeight))
			draw.Draw(part, part.Bounds(), input, image.Point{x, y}, draw.Src)
			parts[i*wNum+j] = part
			in++
		}
	}

	// 随机调换两块
	index1 := rand.Intn(num)
	index2 := rand.Intn(num)
	for index1 == index2 {
		index2 = rand.Intn(num)
	}

	parts[index1], parts[index2] = parts[index2], parts[index1]
	index[index1], index[index2] = index[index2], index[index1]

	// 创建新的合并图像
	mergedImg := image.NewRGBA(bounds)

	// 合并图像
	for i := 0; i < hNum; i++ {
		for j := 0; j < wNum; j++ {
			x := j * partWidth
			y := i * partHeight
			draw.Draw(mergedImg, image.Rect(x, y, x+partWidth, y+partHeight), parts[i*4+j], image.Point{0, 0}, draw.Src)
		}
	}

	return mergedImg, index
}
