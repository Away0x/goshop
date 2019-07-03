package constants

import (
	"strings"

	"github.com/labstack/echo/v4"
)

const (
	// RestfulAPIPrefix : route api prefix
	RestfulAPIPrefix = "/api"
)

// NeedResponseJSON 判断该请求是否是一个 api 请求 (即不渲染 html，而是响应 json)
func NeedResponseJSON(c echo.Context) bool {
	path := c.Request().URL.Path
	return strings.HasPrefix(path, RestfulAPIPrefix) || (c.Request().Header.Get("X-Requested-With") != "")
}
