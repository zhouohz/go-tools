package slide

import (
	"github.com/zhouohz/go-tools/core/file"
	image2 "github.com/zhouohz/go-tools/core/image"
	"github.com/zhouohz/go-tools/core/util/random"
	"image"
)

var innerBlock []Block

func init() {

	name, _ := file.ListFileName(blockDir, true, false)
	innerBlock = dictBlock(name)
}

const (
	Active   = "active.png"
	Fixed    = "fixed.png"
	blockDir = "captcha/resources/slider/block"
)

type Block struct {
	Active image.Image
	Fixed  image.Image
}

func RandGetBlock() Block {
	return innerBlock[random.RandInt(len(innerBlock))]
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
