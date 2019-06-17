package middleware

import (
	"github.com/labstack/echo/v4"
)

// Guest 未登录用户才可访问
func Guest(h echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		return h(c)
	}
}
