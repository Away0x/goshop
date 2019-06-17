package routes

import (
	"echo_shop/pkg/constants"
	customContext "echo_shop/routes/middleware"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Register 注册路由
func Register(e *echo.Echo) {
	// 自定义 context
	e.Use(customContext.Context)
	// from value 中携带 constants.MethodOverrideFromFormKeyName 参数
	// 可以指定请求从 POST 重写为其他 (DELETE、PUT、PATCH ...)
	e.Pre(middleware.MethodOverrideWithConfig(middleware.MethodOverrideConfig{
		Getter: middleware.MethodFromForm(constants.MethodOverrideFromFormKeyName),
	}))

	registerWeb(e)
}
