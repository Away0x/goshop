package routes

import (
	"echo_shop/config"
	"echo_shop/pkg/constants"
	"echo_shop/pkg/errno"
	customContext "echo_shop/routes/middleware"
	mymiddleware "echo_shop/routes/middleware"
	"net/http"

	"echo_shop/app/controllers/sd"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// SpecialHandlers -
type SpecialHandlers struct {
	// 错误处理的 handler
	ErrorHandler func(echo.Context, *errno.Errno) error
}

const (
	restfulAPIPrefix = "/api"
)

// Register 注册路由
func Register(e *echo.Echo) *SpecialHandlers {
	// 自定义 context
	e.Use(customContext.Context)

	// recover
	// e.Use(middleware.Recover())
	e.Use(mymiddleware.Recover())

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

	// 服务器健康自检
	sdRouter := e.Group("/sd")
	{
		sdRouter.GET("/health", sd.HealthCheck).Name = "sd.health"
		sdRouter.GET("/disk", sd.DiskCheck).Name = "sd.disk"
		sdRouter.GET("/cpu", sd.CPUCheck).Name = "sd.cpu"
		sdRouter.GET("/ram", sd.RAMCheck).Name = "sd.ram"
	}

	// 注册 web routes
	registerWeb(e)

	// 注册 api routes
	registerAPI(e, restfulAPIPrefix)

	// 注册 error handler
	echo.NotFoundHandler = notFoundHandler
	echo.MethodNotAllowedHandler = notFoundHandler

	return &SpecialHandlers{
		ErrorHandler: errorHandler,
	}
}
