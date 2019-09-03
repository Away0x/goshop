package middleware

import (
	"echo_shop/app/context"
	"echo_shop/pkg/constants"
	"echo_shop/pkg/errno"

	"github.com/labstack/echo/v4"
)

// SessionAuth 登录用户才可访问 (会验证用户是否激活)
func SessionAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cc := context.NewAppContext(c)
		user := cc.SessionCurrentUser()
		if user == nil {
			if constants.IsAPIRequest(c) {
				return errno.LoginRequiredErr
			}
			return cc.RedirectToLoginPage()
		}

		// 验证用户是否激活
		if !user.IsActivated() {
			if constants.IsAPIRequest(c) {
				return errno.UserActivateErr
			}
			return cc.RedirectToUserVerificationPage()
		}

		return next(c)

	}
}

// SessionAuthAndNoCheckActived 登录用户才可访问 (不会验证用户是否激活)
func SessionAuthAndNoCheckActived(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cc := context.NewAppContext(c)
		if !cc.SessionIsLogin() {
			if constants.IsAPIRequest(c) {
				return errno.LoginRequiredErr
			}
			return cc.RedirectToLoginPage()
		}

		return next(c)
	}
}

// SessionAuthIfLoginCheckActived 如果用户登录需要验证用户是否激活
func SessionAuthIfLoginCheckActived(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cc := context.NewAppContext(c)
		user := cc.SessionCurrentUser()

		if user != nil && !user.IsActivated() {
			if constants.IsAPIRequest(c) {
				return errno.UserActivateErr
			}
			return cc.RedirectToUserVerificationPage()
		}

		return next(c)
	}
}

// SessionGuest 未登录用户才可访问
func SessionGuest(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cc := context.NewAppContext(c)
		if !cc.SessionIsLogin() {
			return next(c)
		}

		return cc.RedirectToHomePage()
	}
}
