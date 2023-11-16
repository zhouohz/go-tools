package captcha

import (
	"github.com/zhouohz/go-tools/core/file"
	image2 "github.com/zhouohz/go-tools/core/image"
	"github.com/zhouohz/go-tools/core/util/random"
	"image"
)

var innerBg []image.Image

const bgDir = "captcha/resources/bg"

func init() {
	suffix, _ := file.FileListBySuffix(bgDir, ".jpg", true, false, 0)
	innerBg = loadBgByNames(suffix)
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

func RandGetBg() image.Image {
	return innerBg[random.RandInt(len(innerBg))]
}
