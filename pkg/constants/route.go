package constants

import (
	"strings"

	"github.com/labstack/echo/v4"
)

const (
	// StaticURLPrefix : 项目静态文件 url prefix
	StaticURLPrefix = "/public"
	// RestfulAPIPrefix : route api prefix
	RestfulAPIPrefix = "/api"
	// AdminWebPrefix : 后台管理页面 prefix
	AdminWebPrefix = "/admin"
)

// NeedResponseJSON 判断该请求是否是一个 api 请求 (即不渲染 html，而是响应 json)
func NeedResponseJSON(c echo.Context) bool {
	path := c.Request().URL.Path
	return strings.HasPrefix(path, RestfulAPIPrefix) || (c.Request().Header.Get("X-Requested-With") != "")
}
