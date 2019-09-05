package routes

import (
	"echo_shop/app/context"
	"echo_shop/bootstrap/config"
	"echo_shop/pkg/constants"
	"echo_shop/pkg/echohelper/handler"
	"echo_shop/pkg/echohelper/middleware"
	"net/http"
	"strings"

	"echo_shop/pkg/captcha"

	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
)

const (
	// CaptchaURL 验证码的 url
	CaptchaURL = "/captcha"
)

// Register 注册路由
func Register(e *echo.Echo) {
	// BEGIN ----------------------------------------- 全局中间件注册 -----------------------------------------
	// 自定义 context
	e.Use(context.WrapContextMiddleware)

	// recover
	e.Use(middleware.Recover())

	if config.IsDev() {
		// log
		e.Use(middleware.Logger(constants.StaticURLPrefix))
	}

	// gzip
	if config.Bool("APP.GZIP") {
		e.Use(echomiddleware.GzipWithConfig(echomiddleware.GzipConfig{
			Skipper: func(c echo.Context) bool {
				url := c.Request().URL.Path
				// 验证码不能 gzip
				if strings.HasPrefix(url, CaptchaURL) {
					return true
				}

				return false
			},
		}))
	}

	// from value 中携带 constants.MethodOverrideFromFormKeyName 参数
	// 可以指定请求从 POST 重写为其他 (DELETE、PUT、PATCH ...)
	e.Pre(echomiddleware.MethodOverrideWithConfig(echomiddleware.MethodOverrideConfig{
		Getter: echomiddleware.MethodFromForm(constants.MethodOverrideFromFormKeyName),
	}))

	// 去除 url 尾部 /
	e.Pre(echomiddleware.RemoveTrailingSlashWithConfig(echomiddleware.TrailingSlashConfig{
		RedirectCode: http.StatusMovedPermanently,
	}))
	// END ----------------------------------------- 全局中间件注册 -----------------------------------------

	// BEGIN ----------------------------------------- 全局路由 & handler -----------------------------------------
	// 错误处理的路由
	registerError(e)

	// 项目静态文件配置
	e.Static(constants.StaticURLPrefix, config.String("APP.PUBLIC_DIR"))
	e.File("/favicon.ico", config.String("APP.PUBLIC_DIR")+"/favicon.ico")

	// 服务器健康自检
	handler.RegisterSDHandlers(e)

	// 验证码
	context.RegisterHandler(e.GET, CaptchaURL+"/:id", func(c *context.AppContext) error {
		return captcha.Handler(c)
	}).Name = "captcha"
	// END ----------------------------------------- 全局 handler -----------------------------------------

	// 注册 web routes
	registerWeb(e)

	// 注册 api routes
	registerAPI(e)

	// 注册后台管理 routes
	registeAdmin(e)
}
