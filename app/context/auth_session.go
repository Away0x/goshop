package context

// auth session 相关

import (
	"echo_shop/app/auth/session"
	"echo_shop/app/models"
)

// SessionCurrentUser 获取当前 user model
func (a *AppContext) SessionCurrentUser() *models.User {
	if user, ok := session.GetCurrentUserFromSession(a.Context); ok {
		return user
	}

	return nil
}

// SessionIsLogin 判断用户是否已登录
func (a *AppContext) SessionIsLogin() bool {
	_, ok := session.GetCurrentUserFromSession(a.Context)
	return ok
}

// SessionLogin 用户登录
func (a *AppContext) SessionLogin(u *models.User) {
	session.Login(a, u)
}

// SessionLogout 用户退出
func (a *AppContext) SessionLogout() {
	session.Logout(a)
}
