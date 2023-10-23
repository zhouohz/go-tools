package slide

import (
	"github.com/zhouohz/go-tools/core/util/random"
	"math"
)

type rg struct {
	max, min int
}

// generatePoint 生成坐标点
func generatePoint(bg, crop *Image) (int, int) {
	//XR := getPointRange(bg.Width, crop.Width)
	//YR := getPointRange(bg.Height, crop.Height)
	//fmt.Println(XR.min)
	//fmt.Println(XR.max)
	//fmt.Println(YR.min)
	//fmt.Println(YR.max)
	//return random.RandomIntRange(XR.min, XR.max), random.RandomIntRange

	widthDifference := bg.Width - crop.Width
	heightDifference := bg.Height - crop.Height

	x, y := 0, 0

	if widthDifference <= 0 {
		x = 5
	} else {
		x = random.RandomIntRange(100, widthDifference-100)
	}
	if heightDifference <= 0 {
		y = 5
	} else {
		y = random.RandomIntRange(5, heightDifference)
	}
	return x, y
}

func getPointRange(bgLen, blockLen int) *rg {
	size := int(math.Ceil(float64(blockLen) / 2.0))
	return &rg{
		min: size,
		max: bgLen - size,
	}
}
