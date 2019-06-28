package routes

import (
	"echo_shop/pkg/errno"
	"echo_shop/routes/middleware"
	"time"

	"github.com/labstack/echo/v4"
)

// NewRateLimiter 限制指定时间内最多多少个请求
// 例: NewRateLimiter(1*time.Minute, 6), 表示 1 分钟最多 6 次请求
func NewRateLimiter(duration time.Duration, n int) echo.MiddlewareFunc {
	config := middleware.RateLimiterConfig{
		Duration: duration,
		MaxCount: n,
		ErrorHandler: func(c echo.Context) error {
			if needResponseJSON(c) {
				return errno.TooManyRequestsErr
			}

			return errno.TooManyRequestsErr.
				SetSummary("太多请求").
				SetMessage("很抱歉！您向我们的服务器发出太多请求了。").
				HTML()
		},
	}

	return middleware.RateLimiter(config)
}
