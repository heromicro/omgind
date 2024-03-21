package wirex

import (
	"github.com/heromicro/omgind/pkg/global"
	"github.com/heromicro/omgind/pkg/mw/rdb"
	"github.com/heromicro/omgind/pkg/vcode"
)

func InitVcode(rdb *rdb.Redis) *vcode.Vcode {

	cfg := global.CFG.Captcha

	vc := vcode.New(rdb, cfg)
	return vc
}
