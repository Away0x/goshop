package echoinit

import (
	"echo_shop/pkg/errno"
	"echo_shop/routes"
	"net/http"

	"github.com/labstack/echo/v4"
)

// SetupError 项目统一错误处理
func SetupError(e *echo.Echo, sh *routes.SpecialHandlers) {
	e.HTTPErrorHandler = func(err error, c echo.Context) {
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
				err = sh.ErrorHandler(c, data)
			}
			if err != nil {
				e.Logger.Error(err)
			}
		}
	}
}
