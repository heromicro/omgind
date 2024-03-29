package vcode

import (
	"errors"
	"fmt"
	"io"
	"time"

	"github.com/mojocn/base64Captcha"
	"github.com/redis/go-redis/v9"

	"github.com/heromicro/omgind/pkg/config/option"
	"github.com/heromicro/omgind/pkg/helper/str"
	"github.com/heromicro/omgind/pkg/helper/uid/uuid"
	"github.com/heromicro/omgind/pkg/mw/rdb"
	"github.com/heromicro/omgind/pkg/vcode/store"
)

var (
	ErrNotFound = errors.New("captcha: id not found")
)

type Vcode struct {
	redisCli redis.UniversalClient
	// rdb    *rdb.Redis
	cpch   *base64Captcha.Captcha
	store  base64Captcha.Store
	driver base64Captcha.Driver
	source string
}

func New(rdb *rdb.Redis, cfg option.CaptchaConfig) *Vcode {

	fmt.Printf(" === captach config %+v \n", cfg)

	driver := base64Captcha.NewDriverString(cfg.Height, cfg.Width, cfg.NoiseCount, cfg.ShowLineOptions, cfg.Length, cfg.Source, cfg.BgColor, nil, cfg.Fonts)

	if cfg.Store == "redis" {
		storer := store.NewRedisStore(rdb, time.Minute*time.Duration(cfg.Duration), cfg.RedisPrefix)

		cpch := base64Captcha.NewCaptcha(driver.ConvertFonts(), storer)
		return &Vcode{
			redisCli: rdb.Client(),
			cpch:     cpch,
			driver:   driver,
			store:    storer,
			source:   cfg.Source,
		}

	} else {
		storer := base64Captcha.NewMemoryStore(base64Captcha.GCLimitNumber, time.Minute*time.Duration(cfg.Duration))
		cpch := base64Captcha.NewCaptcha(driver.ConvertFonts(), storer)
		return &Vcode{
			redisCli: nil,
			cpch:     cpch,
			driver:   driver,
			store:    storer,
			source:   cfg.Source,
		}
	}
}

func (vc *Vcode) NewLen(length int) (string, error) {
	id := uuid.MustString()
	val := str.RandCustom(length, vc.source)
	err := vc.store.Set(id, val)
	return id, err
}

func (vc *Vcode) Reload(id string) bool {
	val := vc.store.Get(id, false)
	if val == "" {
		return false
	}
	vc.store.Set(id, str.RandCustom(len(val), vc.source))
	return true
}

func (vc *Vcode) GenerateImage(w io.Writer, id string) error {

	val := vc.store.Get(id, false)
	if val == "" {
		return errors.New("cap")
	}
	item, err := vc.driver.DrawCaptcha(val)
	if err != nil {
		return err
	}

	_, err = item.WriteTo(w)
	if err != nil {
		return err
	}
	return nil
}

func (vc *Vcode) GenerateBase64(id string) (string, error) {

	val := vc.store.Get(id, false)
	item, err := vc.driver.DrawCaptcha(val)
	if err != nil {
		return "", err
	}
	bs64 := item.EncodeB64string()
	return bs64, nil
}

func (vc *Vcode) Verify(id, answer string, clear bool) bool {
	return vc.store.Verify(id, answer, clear)
}
