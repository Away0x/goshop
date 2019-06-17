package routes

import (
	"echo_shop/config"
	"echo_shop/pkg/constants"
	"echo_shop/pkg/errno"
	customContext "echo_shop/routes/middleware"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const (
	restfulAPIPrefix = "/api"
)

var (
	notFoundHandler = func(c echo.Context) error {
		if strings.HasPrefix(c.Path(), restfulAPIPrefix) {
			return errno.NotFoundErr
		} else {
			return errno.NotFoundErr.SetMessage("很抱歉！您浏览的页面不存在。").RenderNoContent()
		}
	}
)

// Register 注册路由
func Register(e *echo.Echo) {
	// 自定义 context
	e.Use(customContext.Context)

	// recover
	e.Use(middleware.Recover())

	if config.IsDev() {
		// log
		e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
			Format: "${status}  ${method}   ${uri}\n",
		}))
	} else {
		// gzip
		e.Use(middleware.Gzip())
	}

	// from value 中携带 constants.MethodOverrideFromFormKeyName 参数
	// 可以指定请求从 POST 重写为其他 (DELETE、PUT、PATCH ...)
	e.Pre(middleware.MethodOverrideWithConfig(middleware.MethodOverrideConfig{
		Getter: middleware.MethodFromForm(constants.MethodOverrideFromFormKeyName),
	}))

	// 去除 url 尾部 /
	e.Pre(middleware.RemoveTrailingSlashWithConfig(middleware.TrailingSlashConfig{
		RedirectCode: http.StatusMovedPermanently,
	}))

	// 注册 web routes
	registerWeb(e)

	// 注册 api routes

	// 注册 error handler
	echo.NotFoundHandler = notFoundHandler
	echo.MethodNotAllowedHandler = notFoundHandler
}
