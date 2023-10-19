package slide

import (
	"fmt"
	"math"
	"tools/core/util/random"
)

type rg struct {
	max, min int
}

// generatePoint 生成坐标点
func generatePoint(bg, crop *Image) (int, int) {
	XR := getPointRange(bg.Width, crop.Width)
	YR := getPointRange(bg.Height, crop.Height)
	fmt.Println(XR.min)
	fmt.Println(XR.max)
	fmt.Println(YR.min)
	fmt.Println(YR.max)
	return random.RandomIntRange(XR.min, XR.max), random.RandomIntRange(YR.min, YR.max)
}

func getPointRange(bgLen, blockLen int) *rg {
	size := int(math.Ceil(float64(blockLen) / 2.0))
	return &rg{
		min: size,
		max: bgLen - size,
	}
}
