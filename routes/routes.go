package routes

import (
	"echo_shop/app/context"
	"echo_shop/bootstrap/config"
	"echo_shop/pkg/constants"
	"echo_shop/pkg/errno"
	mymiddleware "echo_shop/routes/middleware"
	"net/http"
	"strings"

	"echo_shop/app/controllers/sd"

	"echo_shop/pkg/captcha"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/lexkong/log"
)

// Register 注册路由
func Register(e *echo.Echo) {
	// 自定义 context
	e.Use(context.WrapContextMiddleware)

	// recover
	e.Use(mymiddleware.Recover())

	if config.IsDev() {
		// log
		e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
			Skipper: func(c echo.Context) bool {
				url := c.Request().URL.Path
				// 静态文件 url 不 log
				isStaticURL := strings.HasPrefix(url, constants.StaticURLPrefix)
				isFavicon := url == "/favicon.ico"
				return isStaticURL || isFavicon
			},
			Format: "${status}  ${method}   ${uri}\n",
		}))
	}

	// gzip
	if config.Bool("APP.GZIP") {
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
	e.Static(constants.StaticURLPrefix, config.String("APP.PUBLIC_DIR"))
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

	// 注册后台管理 routes
	registeAdmin(e, constants.AdminWebPrefix)

	// 注册 error handler
	echo.NotFoundHandler = notFoundHandler
	echo.MethodNotAllowedHandler = notFoundHandler
	// handler 返回的 error 的处理函数
	e.HTTPErrorHandler = httpErrorHandler
}

// 未找到路由时的 handler
func notFoundHandler(c echo.Context) error {
	if constants.NeedResponseJSON(c) {
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

// 错误处理
func resolveError(c echo.Context, e *errno.Errno) error {
	if (e.Detail != nil) && (e.Detail.Type == errno.RenderDetailTypeHTML) && (e.Detail.Template != "") {
		return c.Render(e.HTTPCode, e.Detail.Template, e.Detail.Content)
	}

	cc := context.NewAppContext(c)
	return cc.RenderErrorJSON(e)
}
