package limit

import "time"

// ImageCaptchaTrack 结构体
type ImageCaptchaTrack struct {
	BgImageWidth      int         `json:"bgImageWidth"`
	BgImageHeight     int         `json:"bgImageHeight"`
	SliderImageWidth  int         `json:"sliderImageWidth"`
	SliderImageHeight int         `json:"sliderImageHeight"`
	StartSlidingTime  time.Time   `json:"startSlidingTime"`
	EndSlidingTime    time.Time   `json:"endSlidingTime"`
	TrackList         []Track     `json:"trackList"`
	Data              interface{} `json:"data"`
}

// Track 结构体
type Track struct {
	X    int    `json:"x"`
	Y    int    `json:"y"`
	T    int    `json:"t"`
	Type string `json:"type"`
}

// EntSlidingTime 兼容方法
func (i *ImageCaptchaTrack) EntSlidingTime() time.Time {
	return i.EndSlidingTime
}

func (i *ImageCaptchaTrack) SetEntSlidingTime(entSlidingTime time.Time) {
	i.EndSlidingTime = entSlidingTime
}
