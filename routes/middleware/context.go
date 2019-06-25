package middleware

import (
	"echo_shop/app/context"

	"github.com/labstack/echo/v4"
)

// Context -
func Context(h echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cc := &context.AppContext{
			Context: c,
		}

		return h(cc)
	}
}
