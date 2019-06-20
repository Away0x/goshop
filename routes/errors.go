package routes

import (
	"echo_shop/pkg/errno"
	"strings"

	"github.com/labstack/echo/v4"
)

// 未找到路由时的 handler
func notFoundHandler(c echo.Context) error {
	if strings.HasPrefix(c.Path(), restfulAPIPrefix) {
		return errno.NotFoundErr
	}

	return errno.NotFoundErr.SetMessage("很抱歉！您浏览的页面不存在。").RenderNoContent()
}

// 统一错误处理 handler
func errorHandler(c echo.Context, e *errno.Errno) error {
	if e.RenderTpl != "" {
		if d, ok := e.Errors.(map[string]interface{}); ok {
			return c.Render(e.HTTPCode, e.RenderTpl, d)
		}

		return c.Render(e.HTTPCode, e.RenderTpl, map[string]interface{}{})
	}

	return c.JSON(e.HTTPCode, e)
}
