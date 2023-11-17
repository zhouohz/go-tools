package config

import "image/color"

// WatermarkConfig 水印设置
type WatermarkConfig struct {
	Enabled  bool       `yaml:"enabled"`
	FontSize int        `yaml:"fontSize"`
	Color    color.RGBA `yaml:"color"`
	Text     string     `yaml:"text"`
}

type Config struct {
	Watermark      *WatermarkConfig `yaml:"watermark"`
	CacheExpireSec int              `yaml:"cacheExpireSec"`
	// 项目的绝对路径: 图片、字体等
	ResourcePath string `yaml:"resourcePath"`
}
