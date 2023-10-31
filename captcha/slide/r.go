package slide

import (
	image2 "github.com/zhouohz/go-tools/core/image"
	"image"
)

var InnerBg = loadBgByNames([]string{
	"captcha/resources/slider/bg/1.jpg",
})

var InnerBLock = dictBlock([]string{
	"captcha/resources/slider/block/1",
})

type Block struct {
	Active image.Image
	Fixed  image.Image
}

func dictBlock(dir []string) []Block {
	imgs := make([]Block, len(dir))
	for i := range dir {
		imgs[i].Active = loadBgByName(dir[i] + "/" + Active)
		imgs[i].Fixed = loadBgByName(dir[i] + "/" + Fixed)
	}

	return imgs
}

func loadBgByNames(assetNames []string) []image.Image {
	fonts := make([]image.Image, 0)
	for _, assetName := range assetNames {
		f := loadBgByName(assetName)
		fonts = append(fonts, f)
	}
	return fonts
}

func loadBgByName(name string) image.Image {
	img, err := image2.ParseImage(name)
	if err != nil {
		panic(err)
	}

	return img
}
