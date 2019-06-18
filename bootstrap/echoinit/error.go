package echoinit

import (
	"echo_shop/pkg/errno"
	"net/http"

	"github.com/labstack/echo/v4"
)

// SetupError 项目统一错误处理
func SetupError(e *echo.Echo) {
	e.HTTPErrorHandler = func(err error, c echo.Context) {
		var (
			code       = http.StatusInternalServerError
			data       *errno.Errno
			render     string
			renderData map[string]interface{}
		)

		switch typed := err.(type) {
		// http error
		case *echo.HTTPError:
			code = typed.Code
			data = errno.UnknowErr.SetErrors(typed.Message.(string))
			// 自定义错误
		case *errno.Errno:
			code = typed.HTTPCode
			data = typed
			if typed.RenderTpl != "" {
				render = typed.RenderTpl
				renderData = typed.Errors.(map[string]interface{})
			}
		default:
			data = errno.UnknowErr.SetErrors(typed.Error())
		}

		// Send response
		if !c.Response().Committed {
			if c.Request().Method == http.MethodHead {
				err = c.NoContent(code)
			} else {
				if render != "" {
					err = c.Render(code, render, renderData)
				} else {
					err = c.JSON(code, data)
				}
			}
			if err != nil {
				e.Logger.Error(err)
			}
		}
	}
}
