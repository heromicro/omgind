package app

import (
	"github.com/go-redis/redis/v7"
	"github.com/heromicro/omgind/pkg/global"
	"github.com/heromicro/omgind/pkg/vcode"
)

func InitVcode(cli redis.Cmdable) *vcode.Vcode {

	cfg := global.CFG.Captcha

	vc := vcode.New(cli, cfg)
	return vc
}
