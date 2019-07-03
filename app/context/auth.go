package context

import (
	"echo_shop/app/auth"
	"echo_shop/app/models"
	"echo_shop/pkg/constants"
	"echo_shop/pkg/errno"
)

// CurrentUser 获取当前 user model
func (a *AppContext) CurrentUser() *models.User {
	if user, ok := auth.GetCurrentUserFromSession(a.Context); ok {
		return user
	}

	return nil
}

// IsLogin 判断用户是否已登录
func (a *AppContext) IsLogin() bool {
	_, ok := auth.GetCurrentUserFromSession(a.Context)
	return ok
}

// Login 用户登录
func (a *AppContext) Login(u *models.User) {
	auth.Login(a, u)
}

// Logout 用户退出
func (a *AppContext) Logout() {
	auth.Logout(a)
}

// AuthCheck 权限检查
func (a *AppContext) AuthCheck(status bool) *errno.Errno {
	if status {
		return nil
	}

	if constants.NeedResponseJSON(a) {
		return errno.UnauthorizedErr.JSON()
	}

	return errno.UnauthorizedErr.HTML()
}

// AuthCheckByFunc 权限检查
func (a *AppContext) AuthCheckByFunc(fn func() bool) *errno.Errno {
	return a.AuthCheck(fn())
}
