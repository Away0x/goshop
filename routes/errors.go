package routes

import (
	"echo_shop/config"
	"echo_shop/pkg/errno"
	"strings"

	"github.com/labstack/echo/v4"
)

// 未找到路由时的 handler
func notFoundHandler(c echo.Context) error {
	if strings.HasPrefix(c.Path(), restfulAPIPrefix) {
		return errno.NotFoundErr
	}

	return errno.NotFoundErr.SetMessage("很抱歉！您浏览的页面不存在。").HTML()
}

// 统一错误处理 handler
func errorHandler(c echo.Context, e *errno.Errno) error {
	if (e.Detail != nil) && (e.Detail.Type == errno.RenderDetailTypeHTML) && (e.Detail.Template != "") {
		return c.Render(e.HTTPCode, e.Detail.Template, e.Detail.Content)
	}

	// 隐藏错误详情 (默认开启)
	if !config.Bool("APP.SHOW_ERROR_DETAIL") {
		e = e.HideErrorDetail()
	}

	return c.JSON(e.HTTPCode, e)
}
