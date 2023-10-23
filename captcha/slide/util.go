package slide

import (
	"github.com/zhouohz/go-tools/core/util/random"
	"image/png"
	"math/rand"
	"os"
)

var bgSimple = LoadBgByNames([]string{
	"res/bg/1.png",
})

var blockSimple = LoadBgByNames([]string{
	"res/block/1.png",
})

func LoadBgByNames(assetNames []string) []*Image {
	fonts := make([]*Image, 0)
	for _, assetName := range assetNames {
		f := LoadBgByName(assetName)
		fonts = append(fonts, f)
	}
	return fonts
}

func LoadBgByName(name string) *Image {
	open, err := os.Open(name)
	if err != nil {
		panic(err)
	}

	img, err := png.Decode(open)
	if err != nil {
		panic(err)
	}

	return &Image{
		Height: img.Bounds().Dy(),
		Width:  img.Bounds().Dx(),
		Src:    name,
	}
}

// RandBgFrom 选择随机的背景
func RandBgFrom(fonts []*Image) *Image {
	fontCount := len(fonts)
	index := rand.Intn(fontCount)
	return fonts[index]
}

func GetBackgroundImage() *Image {
	max := len(bgSimple) - 1
	if max <= 0 {
		max = 1
	}
	return bgSimple[random.RandomIntRange(0, max)]
}

func GetBlockImage() *Image {
	max := len(blockSimple) - 1
	if max <= 0 {
		max = 1
	}
	return blockSimple[random.RandomIntRange(0, max)]
}
