package wirex

import (
	"github.com/heromicro/omgind/pkg/auth"
	"github.com/heromicro/omgind/pkg/auth/jwtauth"
	"github.com/heromicro/omgind/pkg/auth/jwtauth/store/buntdb"
	"github.com/heromicro/omgind/pkg/auth/jwtauth/store/redis"
	"github.com/heromicro/omgind/pkg/config"
	"github.com/heromicro/omgind/pkg/mw/rdb"

	jwt "github.com/golang-jwt/jwt/v4"
)

// InitAuth 初始化用户认证
func InitAuth(conf *config.AppConfig, rdb *rdb.Redis) (auth.Auther, func(), error) {
	cfg := conf.JWTAuth

	var opts []jwtauth.Option
	opts = append(opts, jwtauth.SetExpired(cfg.Expired))
	opts = append(opts, jwtauth.SetSigningKey([]byte(cfg.SigningKey)))
	opts = append(opts, jwtauth.SetKeyfunc(func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, auth.ErrInvalidToken
		}
		return []byte(cfg.SigningKey), nil
	}))

	var method jwt.SigningMethod
	switch cfg.SigningMethod {
	case "HS256":
		method = jwt.SigningMethodHS256
	case "HS384":
		method = jwt.SigningMethodHS384
	default:
		method = jwt.SigningMethodHS512
	}
	opts = append(opts, jwtauth.SetSigningMethod(method))

	var store jwtauth.Storer
	switch cfg.Store {
	case "redis":
		rcfg := conf.Redis
		store = redis.NewStore(&redis.Config{
			Addrs:     []string{rcfg.Addr},
			Password:  rcfg.Password,
			DB:        cfg.RedisDB,
			KeyPrefix: cfg.RedisPrefix,
		}, rdb)
	default:
		s, err := buntdb.NewStore(cfg.FilePath)
		if err != nil {
			return nil, nil, err
		}
		store = s
	}

	auth := jwtauth.New(store, opts...)
	cleanFunc := func() {
		auth.Release()
	}
	return auth, cleanFunc, nil
}
