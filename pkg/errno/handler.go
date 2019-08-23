package errno

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type resolveErrorFunc = func(echo.Context, *Errno) error

// HTTPErrorHandler 统一错误处理 handler
func HTTPErrorHandler(resolve resolveErrorFunc) echo.HTTPErrorHandler {
	return func(err error, c echo.Context) {
		var data *Errno

		switch typed := err.(type) {
		case *echo.HTTPError: // http error
			data = WrapEchoHTTPError(typed)
		case *Errno: // 自定义错误
			data = typed
		default:
			data = UnknowErr.SetErrorContent(typed)
		}

		// Send response
		if !c.Response().Committed {
			if c.Request().Method == http.MethodHead {
				err = c.NoContent(data.HTTPCode)
			} else {
				err = resolve(c, data)
			}
			if err != nil {
				log.Printf("errno.HTTPErrorHandler: %s", err)
			}
		}
	}
}
