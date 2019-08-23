package context

import (
	"github.com/labstack/echo/v4"
)

type (
	// Context 拓展 echo.Context
	Context struct {
		echo.Context
	}
)

// NewContext -
func NewContext(c echo.Context) Context {
	return Context{
		Context: c,
	}
}
