package routes

import (
	"echo_shop/app/context"
	"echo_shop/app/models"
	"echo_shop/pkg/constants"
	"echo_shop/pkg/errno"
	"log"
	"net/http"

	"github.com/Away0x/validate"
	"github.com/labstack/echo/v4"
)

// 项目统一错误处理
func registerError(e *echo.Echo) {
	// 注册 error handler
	echo.NotFoundHandler = notFoundHandler
	echo.MethodNotAllowedHandler = notFoundHandler

	// controller/handler return error 时会进入该函数
	e.HTTPErrorHandler = func(err error, c echo.Context) {
		data := transformErrorType(err)

		// Send response
		if !c.Response().Committed {
			if c.Request().Method == http.MethodHead {
				err = c.NoContent(data.HTTPCode)
			} else {
				// 响应错误的处理
				err = responseErrorResolve(c, data)
			}
			if err != nil {
				log.Printf("errno.HTTPErrorHandler: %s", err)
			}
		}
	}
}

// 未找到路由时的 handler
func notFoundHandler(c echo.Context) error {
	if constants.IsAPIRequest(c) {
		return errno.NotFoundErr
	}

	return errno.NotFoundErr.SetMessage("很抱歉！您浏览的页面不存在。").HTML()
}

// 转换各种错误类型为 *errno.Errno
func transformErrorType(err error) *errno.Errno {
	switch typed := err.(type) {
	case *errno.Errno:
		return typed
	// echo http error
	case *echo.HTTPError:
		return errno.WrapEchoHTTPError(typed)
	// 请求参数错误
	case validate.Messages:
		return errno.ParamsErr.SetErrorContent(map[string][]string(typed))
	// model error
	case *models.ModelError:
		return errno.DatabaseErr.SetErrorContent(typed)
	// 其他 error
	default:
		return errno.UnknowErr.SetErrorContent(typed)
	}
}

// 接收到错误时的处理函数
func responseErrorResolve(c echo.Context, e *errno.Errno) error {
	// 响应 html
	if (e.Detail != nil) && (e.Detail.Type == errno.RenderDetailTypeHTML) && (e.Detail.Template != "") {
		return c.Render(e.HTTPCode, e.Detail.Template, e.Detail.Content)
	}

	// 响应 json
	cc := context.NewAppContext(c)
	return cc.RenderErrorJSON(e)
}
