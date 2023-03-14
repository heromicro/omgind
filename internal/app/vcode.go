package app

import (
	"github.com/go-redis/redis/v8"
	"github.com/heromicro/omgind/pkg/global"
	"github.com/heromicro/omgind/pkg/vcode"
)

func InitVcode(cli redis.UniversalClient) *vcode.Vcode {

	cfg := global.CFG.Captcha

	vc := vcode.New(cli, cfg)
	return vc
}
