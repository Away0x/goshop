package middleware

import (
	"echo_shop/app/helpers/context"

	"github.com/labstack/echo/v4"
)

// Context -
func Context(h echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cc := &context.CustomContext{
			Context: c,
		}

		return h(cc)
	}
}
