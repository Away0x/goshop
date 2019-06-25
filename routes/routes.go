package routes

import (
	"echo_shop/app/context"
	"echo_shop/config"
	"echo_shop/pkg/constants"
	"echo_shop/pkg/errno"
	AppContext "echo_shop/routes/middleware"
	mymiddleware "echo_shop/routes/middleware"
	"net/http"
	"strings"

	"echo_shop/app/controllers/sd"

	"echo_shop/pkg/captcha"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type (
	echoRegisterFn = func(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	// SpecialHandlers -
	SpecialHandlers struct {
		// 错误处理的 handler
		ErrorHandler func(echo.Context, *errno.Errno) error
	}
)

const (
	restfulAPIPrefix = "/api"
)

// 使用该函数注册 handler，会包装 echo.Context 为 *context.AppContext
func registerHandler(fn echoRegisterFn, path string, h context.AppHandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return fn(path, func(c echo.Context) error {
		cc, ok := c.(*context.AppContext)
		if !ok {
			cc = &context.AppContext{
				Context: c,
			}

			return h(cc)
		}

		return h(cc)
	}, m...)
}

// 判断该请求是否是一个 api 请求 (即不渲染 html，而是响应 json)
func needResponseJSON(c echo.Context) bool {
	path := c.Request().URL.Path
	return strings.HasPrefix(path, restfulAPIPrefix)
}

// Register 注册路由
func Register(e *echo.Echo) *SpecialHandlers {
	// 自定义 context
	e.Use(AppContext.Context)

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
		registerHandler(sdRouter.GET, "/health", sd.HealthCheck).Name = "sd.health"
		registerHandler(sdRouter.GET, "/disk", sd.DiskCheck).Name = "sd.disk"
		registerHandler(sdRouter.GET, "/cpu", sd.CPUCheck).Name = "sd.cpu"
		registerHandler(sdRouter.GET, "/ram", sd.HealthCheck).Name = "sd.ram"
	}

	// 验证码
	registerHandler(e.GET, "/captcha/:id", func(c *context.AppContext) error {
		return captcha.Handler(c)
	}).Name = "captcha"

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
