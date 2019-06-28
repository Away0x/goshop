package middleware

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/patrickmn/go-cache"
	"golang.org/x/time/rate"
)

// RateLimiterConfig -
type RateLimiterConfig struct {
	// 指定时间 Duration 内超出 MaxCount 会进入 ErrorHandler
	Duration     time.Duration
	MaxCount     int
	ErrorHandler echo.HandlerFunc
}

// RateLimiter 限制指定时间内最多多少个请求，超出会进入错误页面
func RateLimiter(config RateLimiterConfig) echo.MiddlewareFunc {
	// 默认 5m 过期，每 10m 清除缓存中的过期 key
	limiterCache := cache.New(5*time.Minute, 10*time.Minute)

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			k := c.RealIP()

			limiter, ok := limiterCache.Get(k)
			if !ok {
				var expire time.Duration
				// ip 限制 duration 时间内最多 MaxCount 个请求
				limiter, expire = rate.NewLimiter(rate.Every(config.Duration), config.MaxCount), time.Hour
				limiterCache.Set(k, limiter, expire)
			}

			ok = limiter.(*rate.Limiter).Allow()
			if !ok {
				return config.ErrorHandler(c)
			}

			return next(c)
		}
	}
}
