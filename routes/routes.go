package routes

import (
	"echo_shop/app/context"
	"echo_shop/bootstrap/config"
	"echo_shop/pkg/constants"
	"echo_shop/pkg/echohelper/handler"
	"echo_shop/pkg/echohelper/middleware"
	"echo_shop/pkg/errno"
	"net/http"
	"strings"

	"echo_shop/pkg/captcha"

	"github.com/Away0x/validate"
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
	// 注册 error handler
	echo.NotFoundHandler = notFoundHandler
	echo.MethodNotAllowedHandler = notFoundHandler
	// handler 返回的 error 的处理函数
	e.HTTPErrorHandler = errno.HTTPErrorHandler(
		func(err error) *errno.Errno {
			switch typed := err.(type) {
			// 请求参数错误
			case validate.Messages:
				return errno.ParamsErr.SetErrorContent(map[string][]string(typed))
			default:
				return errno.UnknowErr.SetErrorContent(typed)
			}
		},
		func(c echo.Context, e *errno.Errno) error {
			if (e.Detail != nil) && (e.Detail.Type == errno.RenderDetailTypeHTML) && (e.Detail.Template != "") {
				return c.Render(e.HTTPCode, e.Detail.Template, e.Detail.Content)
			}

			cc := context.NewAppContext(c)
			return cc.RenderErrorJSON(e)
		},
	)

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

// 未找到路由时的 handler
func notFoundHandler(c echo.Context) error {
	if constants.NeedResponseJSON(c) {
		return errno.NotFoundErr
	}

	return errno.NotFoundErr.SetMessage("很抱歉！您浏览的页面不存在。").HTML()
}
