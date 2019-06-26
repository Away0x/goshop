package middleware

import (
	"echo_shop/app/context"

	"github.com/labstack/echo/v4"
)

// Guest 未登录用户才可访问
func Guest(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cc := context.NewAppContext(c)
		if !cc.IsLogin() {
			return next(c)
		}

		return cc.RedirectToHomePage()
	}
}
