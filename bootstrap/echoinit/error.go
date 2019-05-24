package echoinit

import (
	"echo_shop/pkg/errno"
	"net/http"

	"github.com/labstack/echo/v4"
)

// SetupError -
func SetupError(e *echo.Echo) {
	e.HTTPErrorHandler = func(err error, c echo.Context) {
		var (
			code       = http.StatusInternalServerError
			data       interface{}
			render     string
			renderData map[string]interface{}
		)

		switch typed := err.(type) {
		case *echo.HTTPError:
			code = typed.Code
			data = map[string]interface{}{
				"msg": typed.Message,
			}
		case *errno.Errno:
			code = typed.HTTPCode
			data = typed
			if typed.RenderTpl != "" {
				render = typed.RenderTpl
				renderData = typed.Errors.(map[string]interface{})
			}
		default:
			data = map[string]interface{}{
				"msg": typed.Error(),
			}
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
