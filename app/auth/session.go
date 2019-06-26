package auth

import (
	"echo_shop/app/models"
	"echo_shop/pkg/session"
	"strconv"

	"github.com/labstack/echo/v4"
)

const (
	// session config
	authSessionKey = "user_id"
	authContextKey = "current_user_by_session"
	// remember me config
	// rememberFormKey    = "remember"    // 表单提交时的 form name
	// rememberCookieName = "remember_me" // 存在 cookie 中的 key name
	// rememberMaxAge     = 88888888      // 过期时间
)

// Login 登录
func Login(c echo.Context, u *models.User) {
	session.Set(c, authSessionKey, u.IDString())
}

// Logout 登出
func Logout(c echo.Context) {
	session.Delete(c, authSessionKey)
}

// GetCurrentUserFromSession 通过 session 获取到 user model
func GetCurrentUserFromSession(c echo.Context) (*models.User, bool) {
	id, ok := fromSessionGetUserID(c)
	if !ok {
		return nil, false
	}

	if user, ok := fromContextGetUser(c); ok {
		if int(user.ID) == id {
			return user, ok
		}
		c.Set(authContextKey, nil)
	}

	user, err := models.GetUser(uint(id))
	if err != nil || user == nil {
		return nil, false
	}

	c.Set(authContextKey, user)
	return user, true
}

// ---------------- private
// 从 session 中获取 user id
func fromSessionGetUserID(c echo.Context) (int, bool) {
	idString := session.Get(c, authSessionKey)
	if idString == "" {
		return 0, false
	}

	id, err := strconv.Atoi(idString)
	if err != nil {
		return 0, false
	}

	return id, true
}

// 从 echo context 中获取 user model
func fromContextGetUser(c echo.Context) (*models.User, bool) {
	contextUser := c.Get(authContextKey)
	if contextUser != nil {
		if user, ok := contextUser.(*models.User); ok {
			return user, true
		}
	}

	return nil, false
}
