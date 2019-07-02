package routes

import (
	"echo_shop/config"
	"echo_shop/pkg/errno"
	"net/http"

	"github.com/lexkong/log"

	"github.com/labstack/echo/v4"
)

// 未找到路由时的 handler
func notFoundHandler(c echo.Context) error {
	if needResponseJSON(c) {
		return errno.NotFoundErr
	}

	return errno.NotFoundErr.SetMessage("很抱歉！您浏览的页面不存在。").HTML()
}

// 统一错误处理 handler
func httpErrorHandler(err error, c echo.Context) {
	var data *errno.Errno

	switch typed := err.(type) {
	case *echo.HTTPError: // http error
		data = errno.WrapEchoHTTPError(typed)
	case *errno.Errno: // 自定义错误
		data = typed
	default:
		data = errno.UnknowErr.SetErrorContent(typed)
	}

	// Send response
	if !c.Response().Committed {
		if c.Request().Method == http.MethodHead {
			err = c.NoContent(data.HTTPCode)
		} else {
			err = resolveError(c, data)
		}
		if err != nil {
			log.Error("httpErrorHandler", err)
		}
	}
}

func resolveError(c echo.Context, e *errno.Errno) error {
	if (e.Detail != nil) && (e.Detail.Type == errno.RenderDetailTypeHTML) && (e.Detail.Template != "") {
		return c.Render(e.HTTPCode, e.Detail.Template, e.Detail.Content)
	}

	// 隐藏错误详情 (默认开启)
	if !config.Bool("APP.SHOW_ERROR_DETAIL") {
		e = e.HideErrorDetail()
	}

	return c.JSON(e.HTTPCode, e)
}
