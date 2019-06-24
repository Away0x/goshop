package auth

import (
	"echo_shop/app/models"
	"echo_shop/pkg/session"

	"github.com/labstack/echo/v4"
)

const (
	authSessionKey     = "user_id"
	rememberFormKey    = "remember"    // 表单提交时的 form name
	rememberCookieName = "remember_me" // 存在 cookie 中的 key name
	rememberMaxAge     = 88888888      // 过期时间
)

// Login 登录
func Login(c echo.Context, u *models.User) {
	session.Set(c, authSessionKey, u.IDString())
}

// Logout 登出
func Logout(c echo.Context) {
	session.Delete(c, authSessionKey)
}
