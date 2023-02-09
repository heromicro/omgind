package middleware

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v7"
	"github.com/go-redis/redis_rate/v8"
	"github.com/heromicro/omgind/internal/app/ginx"
	"github.com/heromicro/omgind/pkg/errors"
	"github.com/heromicro/omgind/pkg/global"
)

// RateLimiterMiddleware 请求频率限制中间件
func RateLimiterMiddleware(skippers ...SkipperFunc) gin.HandlerFunc {
	cfg := global.CFG.RateLimiter
	if !cfg.Enable {
		return EmptyMiddleware()
	}

	rc := global.CFG.Redis
	ring := redis.NewRing(&redis.RingOptions{
		Addrs: map[string]string{
			"server1": rc.Addr,
		},
		Password: rc.Password,
		DB:       cfg.RedisDB,
	})

	limiter := redis_rate.NewLimiter(ring)
	// limiter.Fallback = rate.NewLimiter(rate.Inf, 0)

	return func(c *gin.Context) {
		if SkipHandler(c, skippers...) {
			c.Next()
			return
		}

		userID := ginx.GetUserID(c)
		if userID != "" {
			limit := cfg.Count
			result, err := limiter.Allow(userID, redis_rate.PerSecond(int(cfg.Count)))
			if err != nil {
				h := c.Writer.Header()
				h.Set("X-RateLimit-Limit", strconv.FormatInt(limit, 10))
				h.Set("X-RateLimit-Remaining", strconv.FormatInt(int64(result.Remaining), 10))
				h.Set("X-RateLimit-Delay", strconv.FormatInt(int64(result.RetryAfter), 10))
				ginx.ResError(c, errors.ErrTooManyRequests)
				return
			}
		}

		c.Next()
	}
}
