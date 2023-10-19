package slide

import "fmt"

// Slide captcha config for captcha-engine-slide.
type Slide struct {
	Offset             int //容错偏移量
	bgImage, cropImage *Image
}

type Image struct {
	Height int
	Width  int
	Src    string
}

// NewSlide creates driver
func NewSlide(offset int) *Slide {
	return &Slide{Offset: offset}
}

func Get() {
	image := GetBackgroundImage()
	//image.Src
	block := GetBlockImage()
	fmt.Println(generatePoint(image, block))

}
