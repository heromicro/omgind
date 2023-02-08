package middleware

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/go-redis/redis_rate"
	"github.com/heromicro/omgind/internal/app/ginx"
	"github.com/heromicro/omgind/pkg/errors"
	"github.com/heromicro/omgind/pkg/global"
	"golang.org/x/time/rate"
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
	limiter.Fallback = rate.NewLimiter(rate.Inf, 0)

	return func(c *gin.Context) {
		if SkipHandler(c, skippers...) {
			c.Next()
			return
		}

		userID := ginx.GetUserID(c)
		if userID != "" {
			limit := cfg.Count
			rate, delay, allowed := limiter.AllowMinute(userID, limit)
			if !allowed {
				h := c.Writer.Header()
				h.Set("X-RateLimit-Limit", strconv.FormatInt(limit, 10))
				h.Set("X-RateLimit-Remaining", strconv.FormatInt(limit-rate, 10))
				delaySec := int64(delay / time.Second)
				h.Set("X-RateLimit-Delay", strconv.FormatInt(delaySec, 10))
				ginx.ResError(c, errors.ErrTooManyRequests)
				return
			}
		}

		c.Next()
	}
}
