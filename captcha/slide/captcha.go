package slide

import (
	"fmt"
	image2 "github.com/zhouohz/go-tools/core/image"
	"github.com/zhouohz/go-tools/core/util/random"
	"log"
)

type Store interface {
	// Set sets the digits for the captcha id.
	Set(id string, value string)

	// Get returns stored digits for the captcha id. Clear indicates
	// whether the captcha must be deleted from the store.
	Get(id string, clear bool) string

	//Verify captcha's answer directly
	Verify(id, answer string, clear bool) bool
}

//// Driver captcha interface for captcha engine to to write staff
//type Driver interface {
//	//DrawCaptcha draws binary item
//	DrawCaptcha(content string) (item Item, err error)
//	//GenerateIdQuestionAnswer creates rand id, content and answer
//	GenerateIdQuestionAnswer() (id, q, a string)
//}

//// Item is captcha item interface
//type Item interface {
//	//WriteTo writes to a writer
//	WriteTo(w io.Writer) (n int64, err error)
//	//EncodeB64string encodes as base64 string
//	EncodeB64string() string
//}

func GetCaptcha() (string, string, error) {
	background := fmt.Sprintf("D:\\work\\go-tools\\captcha\\slide\\res\\bg\\%d.jpg", 5)
	active := fmt.Sprintf("D:\\work\\go-tools\\captcha\\slide\\res\\block\\%d\\active.png", 1)
	fixed := fmt.Sprintf("D:\\work\\go-tools\\captcha\\slide\\res\\block\\%d\\fixed.png", 1)
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

	base64, err := image2.ToBase64(i)
	if err != nil {
		return "", "", err
	}

	toBase64, err := image2.ToBase64(backOverlayImage)
	if err != nil {
		return "", "", err
	}

	return base64, toBase64, nil
}
