package middleware

import (
	"echo_shop/app/context"

	"github.com/labstack/echo/v4"
)

// Auth 登录用户才可访问 (会验证用户是否激活)
func Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cc := context.NewAppContext(c)
		user := cc.CurrentUser()
		if user == nil {
			return cc.RedirectToLoginPage()
		}

		// 验证用户是否激活
		if !user.IsActivated() {
			return cc.RedirectToUserVerificationPage()
		}

		return next(c)

	}
}

// AuthAndNoCheckActived 登录用户才可访问 (不会验证用户是否激活)
func AuthAndNoCheckActived(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cc := context.NewAppContext(c)
		if !cc.IsLogin() {
			return cc.RedirectToLoginPage()
		}

		return next(c)
	}
}
