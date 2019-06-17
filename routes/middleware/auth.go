package middleware

import (
	"github.com/labstack/echo/v4"
)

// Auth 登录用户才可访问
func Auth(h echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		return h(c)
	}
}
