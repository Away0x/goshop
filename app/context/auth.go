package context

import (
	"echo_shop/app/auth"
	"echo_shop/app/models"
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
