package context

import (
	"github.com/labstack/echo/v4"
)

type (
	// AppContext : 项目自定义的 context，包装了 echo.Context，提供更多方法
	AppContext struct {
		echo.Context
	}
	// AppHandlerFunc : 项目 handler type
	AppHandlerFunc = func(c *AppContext) error
	// echo 注册路由的函数
	echoRegisterFn = func(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
)

// NewAppContext -
func NewAppContext(c echo.Context) *AppContext {
	return &AppContext{
		Context: c,
	}
}

// RegisterHandler 使用该函数注册 handler，会包装 echo.Context 为 *context.AppContext
func RegisterHandler(fn echoRegisterFn, path string, h AppHandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return fn(path, func(c echo.Context) error {
		cc, ok := c.(*AppContext)
		if !ok {
			cc = NewAppContext(c)
			return h(cc)
		}

		return h(cc)
	}, m...)
}

// WrapContextMiddleware 包装 echo.Context 为 *AppContext 的中间件
func WrapContextMiddleware(h echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cc := NewAppContext(c)
		return h(cc)
	}
}
