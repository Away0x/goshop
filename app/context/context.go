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
)
