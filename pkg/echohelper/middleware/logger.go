package middleware

import (
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Logger log middleware
func Logger(staticURL string) echo.MiddlewareFunc {
	return middleware.LoggerWithConfig(middleware.LoggerConfig{
		Skipper: func(c echo.Context) bool {
			url := c.Request().URL.Path
			// 静态文件 url 不 log
			isStaticURL := strings.HasPrefix(url, staticURL)
			isFavicon := url == "/favicon.ico"
			return isStaticURL || isFavicon
		},
		Format: "${status}  ${method}   ${uri}\n",
	})
}
