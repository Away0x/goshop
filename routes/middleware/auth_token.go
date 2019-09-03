package middleware

import (
	"echo_shop/app/auth/jwt"
	"echo_shop/pkg/errno"

	"github.com/labstack/echo/v4"
)

// TokenAuth 存在 token 并且能够获取到用户，并且用户需要激活才可进入
func TokenAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		_, curUser, err := jwt.GetTokenAndUser(c)
		if err != nil {
			return err
		}

		// 验证用户是否激活
		if !curUser.IsActivated() {
			return errno.UserActivateErr
		}

		return next(c)
	}
}
