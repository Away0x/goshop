package routes

import (
	"echo_shop/app/context"
	"echo_shop/bootstrap/config"
	"echo_shop/pkg/constants"
	mymiddleware "echo_shop/routes/middleware"
	"net/http"

	"echo_shop/app/controllers/sd"

	"echo_shop/pkg/captcha"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Register 注册路由
func Register(e *echo.Echo) {
	// 自定义 context
	e.Use(context.WrapContextMiddleware)

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

	// 项目静态文件配置
	e.Static("/public", config.String("APP.PUBLIC_DIR"))
	e.File("/favicon.ico", config.String("APP.PUBLIC_DIR")+"/favicon.ico")

	// 服务器健康自检
	sdRouter := e.Group("/sd")
	{
		context.RegisterHandler(sdRouter.GET, "/health", sd.HealthCheck).Name = "sd.health"
		context.RegisterHandler(sdRouter.GET, "/disk", sd.DiskCheck).Name = "sd.disk"
		context.RegisterHandler(sdRouter.GET, "/cpu", sd.CPUCheck).Name = "sd.cpu"
		context.RegisterHandler(sdRouter.GET, "/ram", sd.HealthCheck).Name = "sd.ram"
	}

	// 验证码
	context.RegisterHandler(e.GET, "/captcha/:id", func(c *context.AppContext) error {
		return captcha.Handler(c)
	}).Name = "captcha"

	// 注册 web routes
	registerWeb(e)

	// 注册 api routes
	registerAPI(e, constants.RestfulAPIPrefix)

	// 注册 error handler
	echo.NotFoundHandler = notFoundHandler
	echo.MethodNotAllowedHandler = notFoundHandler
	// handler 返回的 error 的处理函数
	e.HTTPErrorHandler = httpErrorHandler

}
