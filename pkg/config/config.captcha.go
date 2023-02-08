package config

import (
	"image/color"
)

// Captcha 图形验证码配置参数
type CaptchaConfig struct {
	Store string

	Length          int
	Width           int
	Height          int
	NoiseCount      int
	ShowLineOptions int //
	Source          string
	BgColor         *color.RGBA //
	Fonts           []string    //
	Duration        int

	RedisDB     int
	RedisPrefix string
}
